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
	playerOne       = `Player 1`
	mplusNormalFont font.Face
	// mpluBigFont     font.Face
)
var playerTwo = `Player 2`
var captured = `Captured: `
var exit = `Exit`
var newGame = `New Game`

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

func draw(screen *ebiten.Image, G *Game) {
	/// Draw background
	screen.Fill(color.RGBA{0xaf, 0xaf, 0xff, 0xff})

	/// Draw goban
	opGoban := &ebiten.DrawImageOptions{}
	opGoban.GeoM.Translate(885, 80)
	opGoban.GeoM.Scale(0.7, 0.7)
	screen.DrawImage(imgGoban, opGoban)

	/// Draw stones
	var y int8
	var x int8
	for y = 0; y < 19; y++ {
		for x = 0; x < 19; x++ {
			coordinate := coordinate{y, x}
			if PositionOccupied(coordinate, &G.goban) == true {
				opStone := &ebiten.DrawImageOptions{}
				opStone.GeoM.Translate((838 + (float64(coordinate.y) * 104.6)), (34 + (float64(coordinate.x) * 104.6)))
				opStone.GeoM.Scale(0.7, 0.7)
				if PositionOccupiedByPlayer(coordinate, &G.goban, G.player) == true {
					screen.DrawImage(imgBlack, opStone)
				} else {
					screen.DrawImage(imgWhite, opStone)
				}
			}
		}
	}

	/// Draw text
	text.Draw(screen, playerOne, mplusNormalFont, 80, 120, color.Black)
	text.Draw(screen, captured, mplusNormalFont, 80, 200, color.Black)
	text.Draw(screen, strconv.Itoa(int(G.capture0)), mplusNormalFont, 340, 200, color.Black)

	text.Draw(screen, playerTwo, mplusNormalFont, 2080, 120, color.White)
	text.Draw(screen, captured, mplusNormalFont, 2080, 200, color.White)
	text.Draw(screen, strconv.Itoa(int(G.capture1)), mplusNormalFont, 2340, 200, color.White)

	text.Draw(screen, exit, mplusNormalFont, 2080, 1300, color.Black)    //red
	text.Draw(screen, newGame, mplusNormalFont, 2080, 1200, color.Black) //red
}

func (G *Game) updateGoban() {
	// fmt.Println("Well og boboin...")///////
	// // coordinate := coordinate{0, 0}//////
	// fmt.Println("Well og boboins...")///////
	// PlaceIfValid(coordinate, G)//////
	PlaceRandomIfValid(G)/////////
	// fmt.Println("Well hi there...")///////
	// // DumpGoban(&G.goban) //////
	// fmt.Println("Well hello there...")//////
}

func (G *Game) UpdateGame(fortytwo int) {////listen for input, update struct
	// g.input.Update()
	// if err := g.board.Update(g.input); err != nil {
	// 	return err
	// }
	// fmt.Printf("Og ogah...s\n")///////
	G.updateGoban()
	// fmt.Printf("Well hello you... %d\n", fortytwo)////////
}

func update(screen *ebiten.Image) error {
	// fmt.Println("Hoo ho ho...")///////
	G.UpdateGame(42)
	// fmt.Println("Hi hi hi...")///////

	if ebiten.IsDrawingSkipped() { /// do we want this (see cheat sheet)?
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) == true { // quit if press escape
		os.Exit(0) ////// rm, just for test. Return win message to GUI
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) == true {
		x, y := ebiten.CursorPosition()
		fmt.Printf("Mouse pressed x:%d y:%d\n", x, y)
	}
	draw(screen, G)
	// time.Sleep(1000 * time.Millisecond) //////////
	return nil
}


func RunEbiten() {
	// DumpGoban(&G.goban) //////
	w, h := ebiten.ScreenSizeInFullscreen()
	ebiten.SetFullscreen(true)
	// ebiten.SetCursorVisible(true)//// helpful?
	if err := ebiten.Run(update, w, h, 1, "Gomoku"); err != nil {
		log.Fatal(err)
	}
}
