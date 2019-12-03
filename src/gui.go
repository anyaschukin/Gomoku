package play

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

// update updates and draws the game state
func update(screen *ebiten.Image) error {
	g.updateGame()
	if ebiten.IsDrawingSkipped() {
		// If game runs slowly, rendering result not adopted
		return nil
	}
	draw(screen, g)
	return nil
}

// runGui launches ebiten.Run which calls update 60 times/second
func runGui() {
	w, _ := ebiten.ScreenSizeInFullscreen()
	windowSize := float64(w) / float64(2560)
	ebiten.SetFullscreen(true)
	if err := ebiten.Run(update, 2560, 1440, windowSize, "Gomoku"); err != nil {
		log.Fatal(err)
	}
}
