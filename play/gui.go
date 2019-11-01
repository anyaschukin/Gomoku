package play

import (
	"fmt"
	"log"
	"os"

	"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

var img_goban *ebiten.Image
var img_black *ebiten.Image
var img_white *ebiten.Image

func init() {
	var err error
	img_goban, _, err = ebitenutil.NewImageFromFile("GUI/img/goban.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img_black, _, err = ebitenutil.NewImageFromFile("GUI/img/black.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img_white, _, err = ebitenutil.NewImageFromFile("GUI/img/white.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

// func (G *game) Update() {

// }

// func (G *game) Draw(G *game) {

// }

// func drawStones(goban *[19][19]position, screen *ebiten.Image) (screenWithStones *ebiten.Image) {
// 	DumpGoban(goban) //////
// 	screenWithStones = screen
// 	return screen
// }

func update_game(screen *ebiten.Image) error {
	// if err := GameLoop(); err != nil {//////////update game state
	// 	return err
	// }

	G := GameLoop(InitializeGame())
	// DumpGoban(&G.goban) //////

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
	op_goban := &ebiten.DrawImageOptions{}
	op_goban.GeoM.Translate(825, 20)
	op_goban.GeoM.Scale(0.7, 0.7)

	op_stone := &ebiten.DrawImageOptions{}
	op_stone.GeoM.Translate(838, 34)
	op_stone.GeoM.Scale(0.7, 0.7)

	op_stone2 := &ebiten.DrawImageOptions{}
	op_stone2.GeoM.Translate((838 + 1045), 34)
	op_stone2.GeoM.Scale(0.7, 0.7)

	op_stone3 := &ebiten.DrawImageOptions{}
	op_stone3.GeoM.Translate(838, (34 + 1045))
	op_stone3.GeoM.Scale(0.7, 0.7)

	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	screen.DrawImage(img_goban, op_goban)

	var y int8
	var x int8
	for y = 0; y < 19; y++ {
		for x = 0; x < 19; x++ {
			coordinate := coordinate{y, x}
			if PositionOccupied(coordinate, &G.goban) == true {
				fmt.Println(coordinate)
				op_stone := &ebiten.DrawImageOptions{}
				op_stone.GeoM.Translate((838 + float64(coordinate.y*104)), (34 + float64(coordinate.x*104)))
				op_stone.GeoM.Scale(0.7, 0.7)
				if PositionOccupiedByPlayer(coordinate, &G.goban, G.player) == true {
					screen.DrawImage(img_black, op_stone)
				} else {
					screen.DrawImage(img_white, op_stone)
				}
			}
		}
	}

	// screen = drawStones()
	screen.DrawImage(img_black, op_stone)
	screen.DrawImage(img_black, op_stone2)
	screen.DrawImage(img_black, op_stone3)

	// ebitenutil.DebugPrint(screen, "Our first game in Ebiten!") //////
	// Draw(screen) ////////// draw new image based on new game state
	return nil
}

func RunEbiten(G *game) {
	// if err := ebiten.Run(update, 1500, 1315, 1, "Gomoku"); err != nil {
	// 	log.Fatal(err)
	// }
	// DumpGoban(&G.goban) //////
	w, h := ebiten.ScreenSizeInFullscreen()
	ebiten.SetFullscreen(true)
	// ebiten.SetCursorVisible(true)//// helpful?
	if err := ebiten.Run(update_game, w, h, 1, "Gomoku"); err != nil {
		log.Fatal(err)
	}
	// if err := ebiten.Run(update, screenHeight, screenWidth, 1, "Gomoku"); err != nil {
	// 	log.Fatal(err)
	// }
}
