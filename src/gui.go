package play //gui

import (
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten"
)

func update(screen *ebiten.Image) error {
	g.updateGame()
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	draw(screen, g)
	// time.Sleep(1 * time.Millisecond) ////////// speed?
	return nil
}

func runEbiten() {
	w, h := ebiten.ScreenSizeInFullscreen()
	ebiten.SetFullscreen(true)
	if err := ebiten.Run(update, w, h, 1, "Gomoku"); err != nil {
		log.Fatal(err)
	}
}
