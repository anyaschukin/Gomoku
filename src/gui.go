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
	// time.Sleep(1 * time.Millisecond) ////////// speed?
	return nil
}

// runGui launches ebiten.Run which calls update 60 times/second
func runGui() {
	w, h := ebiten.ScreenSizeInFullscreen()
	// ebiten.SetFullscreen(true) // toggle Fullscreen
	if err := ebiten.Run(update, w, h, 1, "Gomoku"); err != nil {
		log.Fatal(err)
	}
}
