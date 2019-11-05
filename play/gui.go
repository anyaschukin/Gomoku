package play //gui

import (
	// "fmt"

	"log"

	// "time"

	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func gameLoop(coordinate coordinate, G *Game) {
	validated := PlaceIfValid(coordinate, G)
	if validated == true {
		Capture(coordinate, G)
		CheckWin(coordinate, G)
		SwapPlayers(G)
	}
}

func isPlayerHuman(G *Game) bool {
	if (G.player == false && G.ai0.aiplayer == false) ||
		(G.player == true && G.ai1.aiplayer == false) {
		return true
	}
	return false
}

func (G *Game) UpdateGame() { ////listen for input, update struct
	input(G)
	coordinate := coordinate{-1, -1} /////////
	if G.newGame == true {

	} else if G.won == false {
		if isPlayerHuman(G) == true || isPlayerHotseat(G) == true {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) == true {
				x, y := ebiten.CursorPosition()
				if clickGoban(x, y) == true {
					coordinate.x = int8((float64(x) - (zeroX * scale)) / (positionWidth * scale))
					coordinate.y = int8((float64(y) - (zeroY * scale)) / (positionWidth * scale))
					gameLoop(coordinate, G)
				}
			}
		} else { /////////// ai player
			coordinate = RandomCoordinate() // ai suggest move
			gameLoop(coordinate, G)
		} //////// need to integrate hotseat!!!!!!!
	}
}

func update(screen *ebiten.Image) error {
	G.UpdateGame()
	if ebiten.IsDrawingSkipped() { /// do we want this (see cheat sheet)?
		return nil
	}
	draw(screen, G)
	// time.Sleep(1 * time.Millisecond) //////////
	return nil
}

func RunEbiten() {
	w, h := ebiten.ScreenSizeInFullscreen()
	ebiten.SetFullscreen(true)
	// ebiten.SetCursorVisible(true)//// helpful?
	if err := ebiten.Run(update, w, h, 1, "Gomoku"); err != nil {
		log.Fatal(err)
	}
}
