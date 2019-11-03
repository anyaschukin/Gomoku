package play //gui

import (
	"fmt"
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
	/// Draw player AI or Human
	if G.ai0.aiplayer == true {
		text.Draw(screen, artificial, mplusNormalFont, 80, 200, color.Black)
	} else {
		text.Draw(screen, human, mplusNormalFont, 80, 200, color.Black)
	}
	if G.ai1.aiplayer == true {
		text.Draw(screen, artificial, mplusNormalFont, 2050, 200, color.White)
	} else {
		text.Draw(screen, human, mplusNormalFont, 2050, 200, color.White)
	}
	// text.Draw(screen, playerOne, mplusNormalFont, 80, 120, color.Black)
	// text.Draw(screen, playerTwo, mplusNormalFont, 2050, 120, color.White)


	/// Draw Captured
	text.Draw(screen, captured, mplusNormalFont, 80, 300, color.Black)
	text.Draw(screen, strconv.Itoa(int(G.capture0)), mplusNormalFont, 340, 300, color.Black)

	text.Draw(screen, captured, mplusNormalFont, 2050, 300, color.White)
	text.Draw(screen, strconv.Itoa(int(G.capture1)), mplusNormalFont, 2310, 300, color.White)

	/// Draw Options
	text.Draw(screen, exit, mplusNormalFont, 2050, 1300, color.Black)    //red
	text.Draw(screen, newGame, mplusNormalFont, 2050, 1200, color.Black) //red

	/// Draw Messages
	if G.won == false {
		if G.player == false {
			text.Draw(screen, blackMove, mplusNormalFont, 80, 100, color.Black)
			text.Draw(screen, G.message, mplusNormalFont, 80, 400, color.Black)
		} else {
			text.Draw(screen, G.message, mplusNormalFont, 2050, 400, color.White)
			text.Draw(screen, whiteMove, mplusNormalFont, 2050, 100, color.White)
		}
	} else {
		if G.player == true {
			text.Draw(screen, G.message, mplusNormalFont, 80, 400, color.Black)
		} else {
			text.Draw(screen, G.message, mplusNormalFont, 2050, 400, color.White)
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
				opStone.GeoM.Translate((838 + (float64(coordinate.y) * 104.6)), (34 + (float64(coordinate.x) * 104.6)))
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
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) == true {
		x, y := ebiten.CursorPosition()
		fmt.Printf("Mouse pressed x:%d y:%d\n", x, y)
	}
}

func (G *Game) UpdateGame() {////listen for input, update struct
	input(G)
	if G.won == false {
		validated := false
		coordinate := RandomCoordinate() ///// input from user/ai
		validated = PlaceIfValid(coordinate, G)
		if validated == true {
			Capture(coordinate, G)
			CheckWin(coordinate, G)
			SwapPlayers(G)
		}
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
