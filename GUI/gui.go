package gui

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

// const (
// 	screenWidth  = 1315
// 	screenHeight = 1500
// )

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

// func update(screen *ebiten.Image) error {
// 	if ebiten.IsDrawingSkipped() {
// 		return nil
// 	}

// 	return nil
// }

// func main() {
// 	if err := ebiten.Run(update, 640, 480, 1, "Render an image"); err != nil {
// 		log.Fatal(err)
// 	}
// }

func update(screen *ebiten.Image) error {
	// if err := GameLoop(); err != nil {//////////update game state
	// 	return err
	// }
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
	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	screen.DrawImage(img_goban, op_goban)
	screen.DrawImage(img_black, op_goban)
	// ebitenutil.DebugPrint(screen, "Our first game in Ebiten!") //////
	// Draw(screen) ////////// draw new image based on new game state
	return nil
}

func RunEbiten() {
	// if err := ebiten.Run(update, 1500, 1315, 1, "Gomoku"); err != nil {
	// 	log.Fatal(err)
	// }

	w, h := ebiten.ScreenSizeInFullscreen()
	ebiten.SetFullscreen(true)
	// ebiten.SetCursorVisible(true)//// helpful?
	if err := ebiten.Run(update, w, h, 1, "Gomoku"); err != nil {
		log.Fatal(err)
	}
	// if err := ebiten.Run(update, screenHeight, screenWidth, 1, "Gomoku"); err != nil {
	// 	log.Fatal(err)
	// }
}
