package gomoku

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

// newGame initializes a new game
func newGame() *game {
	g = &game{}
	g.ai0.aiPlayer = true
	g.ai0.depth = 4
	g.ai1.depth = 4
	g.gui.drawLastMove = true
	g.gui.drawWinMove = true
	g.gui.drawCapture = true
	aiSuggestMove(g)
	return g
}

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

// LaunchGui initializes a new game & launches ebiten.Run which calls update 60 times/second
func LaunchGui() {
	g := newGame()
	g.gui.newGame = true
	w, _ := ebiten.ScreenSizeInFullscreen()
	windowSize := float64(w) / float64(2560)
	ebiten.SetFullscreen(true)
	if err := ebiten.Run(update, 2560, 1440, windowSize, "Gomoku"); err != nil {
		log.Fatal(err)
	}
}
