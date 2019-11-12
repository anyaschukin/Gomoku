package play //gui

import (
	// "fmt"
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func gameLoop(coordinate Coordinate, G *Game) {
	validated := PlaceIfValid(coordinate, G)
	if validated == true {
		Capture(coordinate, G)
		CheckWin(coordinate, G)
		G.LastMove = coordinate
		SwapPlayers(G)
		G.Move++
	}
	SuggestMove(G)
}

func (G *Game) UpdateGame() { ////listen for input, update struct
	input(G)
	coordinate := Coordinate{-1, -1} /////////
	if G.NewGame == false && G.Won == false {
		if IsPlayerHuman(G) == true || IsPlayerHotseat(G) == true {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) == true {
				x, y := ebiten.CursorPosition()
				if clickGoban(x, y) == true {
					coordinate.X = int8((float64(x) - (zeroX * scale)) / (positionWidth * scale))
					coordinate.Y = int8((float64(y) - (zeroY * scale)) / (positionWidth * scale))
					gameLoop(coordinate, G)
				}
			}
		} else { /////////// ai Player
			if G.Player == false {
				coordinate = G.Ai0.Suggest
			} else {
				coordinate = G.Ai1.Suggest
			}
			gameLoop(coordinate, G)
		}
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
	// ebiten.SetCursorVisible(true) //// helpful?
	if err := ebiten.Run(update, w, h, 1, "Gomoku"); err != nil {
		log.Fatal(err)
	}
}
