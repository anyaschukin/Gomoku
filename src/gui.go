package play //gui

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

// update updates and draws the game state
func update(screen *ebiten.Image) error {
	g.updateGame()
	if ebiten.IsDrawingSkipped() {
		// If game running slowly, rendering result not adopted
		return nil
	}
	draw(screen, g)
	return nil
}

// runGui launches ebiten.Run which calls update 60 times/second
func runGui() {
	w, h := ebiten.ScreenSizeInFullscreen()
	ebiten.SetFullscreen(true) // toggle Fullscreen
	windowSize := float64(2560 / w)
	if err := ebiten.Run(update, w, h, windowSize, "Gomoku"); err != nil {
		log.Fatal(err)
	}
}
