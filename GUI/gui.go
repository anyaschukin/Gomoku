package gui

import (
	"log"
	"os"
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// const (
// 	screenWidth  = 1315
// 	screenHeight = 1500
// )

func update(screen *ebiten.Image) error {
	// if err := GameLoop(); err != nil {//////////update game state
	// 	return err
	// }
	if ebiten.IsDrawingSkipped() { /// do we want this (see cheat sheet)?
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) == true {// quit if press escape
		os.Exit(0) ////// rm, just for test. Return win message to GUI	
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) == true {
		fmt.Println("Oh hi!")
		x, y := ebiten.CursorPosition()
		fmt.Printf("x:%d y:%d\n", x, y)
	}
	ebitenutil.DebugPrint(screen, "Our first game in Ebiten!") //////
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
