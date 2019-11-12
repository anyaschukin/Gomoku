package play //gui

import (
	// "fmt"
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten"
)

func update(screen *ebiten.Image) error {
	G.updateGame()
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	draw(screen, G)
	// time.Sleep(1 * time.Millisecond) //////////
	return nil
}

func RunEbiten() {
	w, h := ebiten.ScreenSizeInFullscreen()
	ebiten.SetFullscreen(true)
	// ebiten.SetCursorVisible(true) //// helpful?
	if err := ebiten.Run(update, w, h, 1, "Gomoku"); err != nil {
		log.Fatal(err)
	}
}
