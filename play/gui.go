package play //gui

import (
	// "fmt"

	"log"
	"os"
	"strconv"

	// "time"

	"image/color"
	_ "image/png"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
)

var imgGoban *ebiten.Image
var imgBlack *ebiten.Image
var imgRed *ebiten.Image
var imgWhite *ebiten.Image
var imgExit *ebiten.Image
var imgNewGame *ebiten.Image
var imgSelect *ebiten.Image

var (
	captured        = `Captured: `
	mplusNormalFont font.Face
	// mpluBigFont     font.Face
)
var exit = `Exit`
var newGame = `New Game`
var blackMove = `Black to Move`
var whiteMove = `White to Move`
var human = `Human`
var artificial = `AI depth`
var hotseat = `Hotseat`

/// Goban positions
var positionWidth float64 = 104.6
var zeroX float64 = 838 // Left
var zeroY float64 = 34  // Top
var scale float64 = 0.7

/// Exit position
var exitX float64 = 3210
var exitY float64 = 1814

/// New Game position
var newGameX float64 = 3405
var newGameY float64 = 1914
var newGameBlack2 float64 = 5000

var row = 100 // screen indent

func init() {
	/// Init images
	var err error
	imgGoban, _, err = ebitenutil.NewImageFromFile("GUI/img/goban.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgBlack, _, err = ebitenutil.NewImageFromFile("GUI/img/black.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgWhite, _, err = ebitenutil.NewImageFromFile("GUI/img/white.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgRed, _, err = ebitenutil.NewImageFromFile("GUI/img/red.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgExit, _, err = ebitenutil.NewImageFromFile("GUI/img/exit.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgNewGame, _, err = ebitenutil.NewImageFromFile("GUI/img/newGame.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgSelect, _, err = ebitenutil.NewImageFromFile("GUI/img/select.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	/// Init text
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    52,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	// mplusBigFont = truetype.NewFace(tt, &truetype.Options{
	// 	Size:    72,
	// 	DPI:     dpi,
	// 	Hinting: font.HintingFull,
	// })
}

func drawExit(screen *ebiten.Image, G *Game) {
	opExit := &ebiten.DrawImageOptions{}
	opExit.GeoM.Translate(exitX, exitY)
	opExit.GeoM.Scale(scale, scale)
	screen.DrawImage(imgExit, opExit)
}

func drawNewGame(screen *ebiten.Image, G *Game) {
	opNewGame := &ebiten.DrawImageOptions{}
	opNewGame.GeoM.Translate(newGameX, newGameY)
	opNewGame.GeoM.Scale(0.6, 0.6)
	screen.DrawImage(imgNewGame, opNewGame)
}

func drawText(screen *ebiten.Image, G *Game) {
	columnBlack := 80   // screen indent
	columnWhite := 2050 // screen indent
	/// Draw player AI or Human
	if G.ai0.aiplayer == true {
		text.Draw(screen, artificial, mplusNormalFont, columnBlack, row*2, color.Black)
		text.Draw(screen, strconv.Itoa(int(G.ai0.depth)), mplusNormalFont, columnBlack+230, row*2, color.Black)
	} else {
		text.Draw(screen, human, mplusNormalFont, columnBlack, row*2, color.Black)
	}
	if G.ai1.aiplayer == true {
		text.Draw(screen, artificial, mplusNormalFont, columnWhite, row*2, color.White)
		text.Draw(screen, strconv.Itoa(int(G.ai1.depth)), mplusNormalFont, columnWhite+230, row*2, color.White)
	} else {
		text.Draw(screen, human, mplusNormalFont, columnWhite, row*2, color.White)
	}

	/// Draw Captured
	text.Draw(screen, captured, mplusNormalFont, columnBlack, row*3, color.Black)
	text.Draw(screen, strconv.Itoa(int(G.capture0)), mplusNormalFont, 340, row*3, color.Black)

	text.Draw(screen, captured, mplusNormalFont, columnWhite, row*3, color.White)
	text.Draw(screen, strconv.Itoa(int(G.capture1)), mplusNormalFont, 2310, row*3, color.White)

	/// Draw Messages
	if G.won == false {
		if G.player == false {
			text.Draw(screen, blackMove, mplusNormalFont, columnBlack, row, color.Black)
			text.Draw(screen, G.message, mplusNormalFont, columnBlack, row*5, color.Black)
		} else {
			text.Draw(screen, whiteMove, mplusNormalFont, columnWhite, row, color.White)
			text.Draw(screen, G.message, mplusNormalFont, columnWhite, row*5, color.White)
		}
	} else {
		if G.message == "Black Wins!" {
			text.Draw(screen, G.message, mplusNormalFont, columnBlack, row*5, color.Black)
		} else {
			text.Draw(screen, G.message, mplusNormalFont, columnWhite, row*5, color.White)
		}
	}
}

func drawGoban(screen *ebiten.Image, G *Game) {
	opGoban := &ebiten.DrawImageOptions{}
	opGoban.GeoM.Translate(885, 80)
	opGoban.GeoM.Scale(scale, scale)
	screen.DrawImage(imgGoban, opGoban)
}

func drawStones(screen *ebiten.Image, G *Game) {
	var y int8
	var x int8
	for y = 0; y < 19; y++ {
		for x = 0; x < 19; x++ {
			coordinate := coordinate{y, x}
			if PositionOccupied(coordinate, &G.goban) == true {
				opStone := &ebiten.DrawImageOptions{}
				opStone.GeoM.Translate((zeroX + (float64(coordinate.x) * positionWidth)), (zeroY + (float64(coordinate.y) * positionWidth)))
				opStone.GeoM.Scale(scale, scale)
				if PositionOccupiedByPlayer(coordinate, &G.goban, false) == true {
					screen.DrawImage(imgBlack, opStone)
				} else {
					screen.DrawImage(imgWhite, opStone)
				}
			}
		}
	}
}

func drawSelectHuman(screen *ebiten.Image, G *Game, shift float64) {
	opSelect := &ebiten.DrawImageOptions{}
	opSelect.GeoM.Translate(newGameBlack2+shift, 1550)
	opSelect.GeoM.Scale(0.12, 0.12)
	screen.DrawImage(imgSelect, opSelect)
}

func drawSelectAI(screen *ebiten.Image, G *Game, shift float64) {
	depth := G.ai0.depth
	if shift != 0 {
		depth = G.ai1.depth
	}
	opSelect := &ebiten.DrawImageOptions{}
	opSelect.GeoM.Translate(950+shift, 1550+(float64(depth)*833))
	opSelect.GeoM.Scale(0.12, 0.12)
	screen.DrawImage(imgSelect, opSelect)
}

func drawSelectHotseat(screen *ebiten.Image, G *Game, shift float64) {
	opSelect := &ebiten.DrawImageOptions{}
	opSelect.GeoM.Translate(newGameBlack2+shift, 2390)
	opSelect.GeoM.Scale(0.12, 0.12)
	screen.DrawImage(imgSelect, opSelect)
}

func drawSelectPlayer(screen *ebiten.Image, G *Game, player bool) {
	p := G.ai0
	var shift float64
	if player == true {
		p = G.ai1
		shift = 9100
	}
	if p.hotseat == true {
		drawSelectHotseat(screen, G, shift)
		drawSelectHuman(screen, G, shift)
		drawSelectAI(screen, G, shift)
	} else if p.aiplayer == false {
		drawSelectHuman(screen, G, shift)
	} else {
		drawSelectAI(screen, G, shift)
	}
}

func drawSelect(screen *ebiten.Image, G *Game) {
	drawSelectPlayer(screen, G, false)
	drawSelectPlayer(screen, G, true)
}

func drawNewGameOptions(screen *ebiten.Image, G *Game) {
	columnBlack := 140
	columnWhite := 1230
	screen.Fill(color.RGBA{0xaf, 0xaf, 0xff, 0xff}) /// Draw background/////

	opBlack := &ebiten.DrawImageOptions{}
	opBlack.GeoM.Translate(float64(columnBlack)+340, 50)
	screen.DrawImage(imgBlack, opBlack)

	opWhite := &ebiten.DrawImageOptions{}
	opWhite.GeoM.Translate(float64(columnWhite)+340, 50)
	screen.DrawImage(imgWhite, opWhite)

	drawSelect(screen, G)

	for i := 0; i <= 11; i++ {
		text.Draw(screen, artificial, mplusNormalFont, columnBlack, row*(i+2)+50, color.Black)
		text.Draw(screen, strconv.Itoa(i), mplusNormalFont, columnBlack+230, row*(i+2)+50, color.Black)
		text.Draw(screen, artificial, mplusNormalFont, columnWhite, row*(i+2)+50, color.White)
		text.Draw(screen, strconv.Itoa(i), mplusNormalFont, columnWhite+230, row*(i+2)+50, color.White)
	}
	text.Draw(screen, human, mplusNormalFont, columnBlack+520, row*2+50, color.Black)
	text.Draw(screen, hotseat, mplusNormalFont, columnBlack+520, row*3+50, color.Black)
	text.Draw(screen, human, mplusNormalFont, columnWhite+520, row*2+50, color.White)
	text.Draw(screen, hotseat, mplusNormalFont, columnWhite+520, row*3+50, color.White)
}

func draw(screen *ebiten.Image, G *Game) {
	screen.Fill(color.RGBA{0xaf, 0xaf, 0xff, 0xff}) /// Draw background
	if G.newGame == true {
		drawNewGameOptions(screen, G)
	} else {
		drawGoban(screen, G)
		drawStones(screen, G)
		drawText(screen, G)
	}
	drawNewGame(screen, G)
	drawExit(screen, G)
}

func input(G *Game) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) == true { // quit if press escape
		os.Exit(0) ////// is this exiting properly?
	}
	// click exit or new game
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) == true {
		x, y := ebiten.CursorPosition()
		// fmt.Printf("mouse pressed x: %d, y: %d\n", x, y) ////////
		if clickExit(x, y) == true {
			os.Exit(0) ////// is this exiting properly?
		}
		if clickNewGame(x, y) == true {
			if G.newGame == false {
				G.newGame = true
			} else {
				blackTmp := G.ai0
				whiteTmp := G.ai1
				G := NewGame() /////// save new game ai settings!!!!!!
				G.ai0 = blackTmp
				G.ai1 = whiteTmp
			}
		}
		if G.newGame == true {
			if clickHuman0(x, y) == true {
				G.ai0.aiplayer = false
			}
			if clickAI0(x, y) == true {
				G.ai0.aiplayer = true
				G.ai0.depth = uint8((y - 186) / (1201 / 12))
			}
			if clickHotseat0(x, y) == true {
				if G.ai0.hotseat == false {
					G.ai0.hotseat = true
				} else {
					G.ai0.hotseat = false
				}
			}
			if clickHuman1(x, y) == true {
				G.ai1.aiplayer = false
			}
			if clickAI1(x, y) == true {
				G.ai1.aiplayer = true
				G.ai1.depth = uint8((y - 186) / (1201 / 12))
			}
			if clickHotseat1(x, y) == true {
				if G.ai1.hotseat == false {
					G.ai1.hotseat = true
				} else {
					G.ai1.hotseat = false
				}
			}
		}
	}
}

func clickHuman0(x, y int) bool {
	if x > int(newGameBlack2*0.12) && x < 940 &&
		y > int(1550*0.12) && y < 277 {
		return true
	}
	return false
}

func clickHuman1(x, y int) bool {
	if x > int(newGameBlack2*0.12)+1092 && x < 940+1092 &&
		y > int(1550*0.12) && y < 277 {
		return true
	}
	return false
}

func clickHotseat0(x, y int) bool {
	if x > int(newGameBlack2*0.12) && x < 940 &&
		y > 286 && y < 378 {
		return true
	}
	return false
}

func clickHotseat1(x, y int) bool {
	if x > int(newGameBlack2*0.12)+1092 && x < 940+1092 &&
		y > 286 && y < 378 {
		return true
	}
	return false
}

func clickAI0(x, y int) bool {
	if x > int(950*0.12) && x < 454 &&
		y > 186 && y < 1378 {
		return true
	}
	return false
}

func clickAI1(x, y int) bool {
	if x > int(950*0.12)+1092 && x < 454+1092 &&
		y > 186 && y < 1378 {
		return true
	}
	return false
}

func clickNewGame(x, y int) bool {
	if x > int(newGameX*0.6) && x < 2492 &&
		y > int(newGameY*0.6) && y < 1240 {
		return true
	}
	return false
}

func clickExit(x, y int) bool {
	if x > int(exitX*scale) && x < 2492 &&
		y > int(exitY*scale) && y < 1372 {
		return true
	}
	return false
}

func clickGoban(x, y int) bool {
	if x > int(zeroX*scale) && x < int((zeroX*scale)+(positionWidth*float64(19)*scale)) &&
		y > int(zeroY*scale) && y < int((zeroY*scale)+(positionWidth*float64(19)*scale)) {
		return true
	}
	return false
}

func gameLoop(coordinate coordinate, G *Game) {
	validated := PlaceIfValid(coordinate, G)
	if validated == true {
		Capture(coordinate, G)
		CheckWin(coordinate, G)
		SwapPlayers(G)
	}
}

func isPlayerHuman(G *Game) bool {
	if (G.player == false && G.ai0.aiplayer == false) ||
		(G.player == true && G.ai1.aiplayer == false) {
		return true
	}
	return false
}

func (G *Game) UpdateGame() { ////listen for input, update struct
	input(G)
	coordinate := coordinate{-1, -1} /////////
	if G.newGame == true {

	} else if G.won == false {
		if isPlayerHuman(G) == true {
			// if hotseat {
			// draw suggestion
			// }
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) == true {
				x, y := ebiten.CursorPosition()
				if clickGoban(x, y) == true {
					coordinate.x = int8((float64(x) - (zeroX * scale)) / (positionWidth * scale))
					coordinate.y = int8((float64(y) - (zeroY * scale)) / (positionWidth * scale))
					gameLoop(coordinate, G)
				}
			}
		} else { /////////// ai player
			coordinate = RandomCoordinate() // ai suggest move
			gameLoop(coordinate, G)
		} //////// need to integrate hotseat!!!!!!!
	}
}

func update(screen *ebiten.Image) error {
	G.UpdateGame()
	if ebiten.IsDrawingSkipped() { /// do we want this (see cheat sheet)?
		return nil
	}
	draw(screen, G)
	// time.Sleep(1 * time.Millisecond) //////////
	return nil
}

func RunEbiten() {
	w, h := ebiten.ScreenSizeInFullscreen()
	ebiten.SetFullscreen(true)
	// ebiten.SetCursorVisible(true)//// helpful?
	if err := ebiten.Run(update, w, h, 1, "Gomoku"); err != nil {
		log.Fatal(err)
	}
}
