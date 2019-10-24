package gui

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 1315
	screenHeight = 1500
)

func update(screen *ebiten.Image) error {
	// if err := GameLoop(); err != nil {//////////update game state
	// 	return err
	// }
	if ebiten.IsDrawingSkipped() { /// do we want this (see cheat sheet)?
		return nil
	}
	ebitenutil.DebugPrint(screen, "Our first game in Ebiten!") //////
	// Draw(screen) ////////// draw new image based on new game state
	return nil
}

func RunEbiten() {
	// if err := ebiten.Run(update, 1500, 1315, 1, "Gomoku"); err != nil {
	// 	log.Fatal(err)
	// }
	if err := ebiten.Run(update, screenHeight, screenWidth, 1, "Gomoku"); err != nil {
		log.Fatal(err)
	}
}
