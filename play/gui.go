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
var imgWhite *ebiten.Image

var (
	playerOne       = `Player 1`//?
	mplusNormalFont font.Face
	// mpluBigFont     font.Face
)
var playerTwo = `Player 2`//?
var captured = `Captured: `
var exit = `Exit`
var newGame = `New Game`
var blackMove = `Black to Move`
var whiteMove = `White to Move`
var human = `Human player`
var artificial = `AI player`

/// Goban positions
var positionWidth float64 = 104.6
var zeroX float64 = 838	// Left
var zeroY float64 = 34  // Top

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

func drawText(screen *ebiten.Image, G *Game) {
	columnBlack := 80 	// screen indent
	columnWhite := 2050 // screen indent
	row := 100			// screen indent
	/// Draw player AI or Human
	if G.ai0.aiplayer == true {
		text.Draw(screen, artificial, mplusNormalFont, columnBlack, row * 2, color.Black)
	} else {
		text.Draw(screen, human, mplusNormalFont, columnBlack, row * 2, color.Black)
	}
	if G.ai1.aiplayer == true {
		text.Draw(screen, artificial, mplusNormalFont, columnWhite, row * 2, color.White)
	} else {
		text.Draw(screen, human, mplusNormalFont, columnWhite, row * 2, color.White)
	}
	// text.Draw(screen, playerOne, mplusNormalFont, columnBlack, 120, color.Black)
	// text.Draw(screen, playerTwo, mplusNormalFont, columnWhite, 120, color.White)


	/// Draw Captured
	text.Draw(screen, captured, mplusNormalFont, columnBlack, row * 3, color.Black)
	text.Draw(screen, strconv.Itoa(int(G.capture0)), mplusNormalFont, 340, row * 3, color.Black)

	text.Draw(screen, captured, mplusNormalFont, columnWhite, row * 3, color.White)
	text.Draw(screen, strconv.Itoa(int(G.capture1)), mplusNormalFont, 2310, row * 3, color.White)

	/// Draw Options
	text.Draw(screen, exit, mplusNormalFont, columnWhite, row * 13, color.Black)    //red
	text.Draw(screen, newGame, mplusNormalFont, columnWhite, row * 12, color.Black) //red

	/// Draw Messages
	if G.won == false {
		if G.player == false {
			text.Draw(screen, blackMove, mplusNormalFont, columnBlack, row, color.Black)
			text.Draw(screen, G.message, mplusNormalFont, columnBlack, row * 4, color.Black)
		} else {
			text.Draw(screen, G.message, mplusNormalFont, columnWhite, row * 4, color.White)
			text.Draw(screen, whiteMove, mplusNormalFont, columnWhite, row, color.White)
		}
	} else {
		if G.player == true {
			text.Draw(screen, G.message, mplusNormalFont, columnBlack, row * 4, color.Black)
		} else {
			text.Draw(screen, G.message, mplusNormalFont, columnWhite, row * 4, color.White)
		}		
	}
}

func drawGoban(screen *ebiten.Image, G *Game) {
	opGoban := &ebiten.DrawImageOptions{}
	opGoban.GeoM.Translate(885, 80)
	opGoban.GeoM.Scale(0.7, 0.7)
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
				opStone.GeoM.Scale(0.7, 0.7)
				if PositionOccupiedByPlayer(coordinate, &G.goban, false) == true {
					screen.DrawImage(imgBlack, opStone)
				} else {
					screen.DrawImage(imgWhite, opStone)
				}
			}
		}
	}
}

func draw(screen *ebiten.Image, G *Game) {
	screen.Fill(color.RGBA{0xaf, 0xaf, 0xff, 0xff}) 	/// Draw background
	drawGoban(screen, G)
	drawStones(screen, G)
	drawText(screen, G)
}

func input(G *Game) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) == true { // quit if press escape
		os.Exit(0) ////// rm, just for test. Return win message to GUI
	}
	// click exit or new game
}

func clickInGoban(x, y int) bool {
	if x > int(zeroX * 0.7) && x < int((zeroX * 0.7) + (positionWidth * float64(19) * 0.7)) &&
		y > int(zeroY * 0.7) && y < int((zeroY * 0.7) + (positionWidth * float64(19) * 0.7)) {
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

func (G *Game) UpdateGame() {////listen for input, update struct
	input(G)
	coordinate := coordinate{-1, -1}/////////
	if G.won == false {
		if (G.player == false && G.ai0.aiplayer == false) || 						////// human player
			(G.player == true && G.ai1.aiplayer == false) {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) == true {
				x, y := ebiten.CursorPosition()
				if clickInGoban(x, y) == true {
					coordinate.x = int8((float64(x) - (zeroX * 0.7)) / (positionWidth * 0.7))
					coordinate.y = int8((float64(y) - (zeroY * 0.7)) / (positionWidth * 0.7))
					gameLoop(coordinate, G)
				}
			}
		} else { 															/////////// ai player
			coordinate = RandomCoordinate()// ai suggest move
			gameLoop(coordinate, G)
		}//////// need to integrate hotseat!!!!!!!
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
