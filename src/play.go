package play

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// newGame initializes a new game
func newGame() *game {
	g = &game{}
	g.ai0.aiPlayer = true
	g.ai0.depth = 2 // 3
	g.ai1.depth = 2 // 3
	g.gui.drawLastMove = true         //// false for correction (bonus)!!!!
	g.gui.drawWinMove = true          //// false for correction (bonus)!!!!
	g.gui.drawCapture = true //// false for correction (bonus)!!!!
	aiSuggestMove(g)
	return g
}

// swapPlayers prepares for the next move
func swapPlayers(coordinate coordinate, g *game) {
	swapBool(&g.player)
	g.lastMove2 = g.lastMove
	g.lastMove = coordinate
	if g.won == false {
		g.gui.message = ""
	}
	g.move++
}

// gameLoop runs one move
func gameLoop(coordinate coordinate, g *game) {
	g.gui.capturedPositions2 = g.gui.capturedPositions
	g.gui.capturedPositions = nil
	validated := placeIfValid(coordinate, g)
	if validated == true {
		capture(coordinate, g)
		checkWin(coordinate, g)
		swapPlayers(coordinate, g)
		if isPlayerHotseat(g) == true {
			aiSuggestMove(g)
		}
	}
}

// humanLoop listens for a click on the goban runs gameloop with clicked coordinate
func humanLoop(g *game) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) == true {
		x, y := ebiten.CursorPosition()
		if clickGoban(x, y) == true {
			coordinate := coordinate{-1, -1}
			coordinate.x = int8((float64(x) - (gobanX * scale)) / (positionWidth * scale))
			coordinate.y = int8((float64(y) - (gobanY * scale)) / (positionWidth * scale))
			gameLoop(coordinate, g)
		}
	}
}

// aiLoop listens runs gameloop with suggested coordinate
func aiLoop(g *game) {
	coordinate := g.ai0.suggest
	if g.player == true {
		coordinate = g.ai1.suggest
	}
	gameLoop(coordinate, g)
	aiSuggestMove(g)
}

// updateGame listens for input, and runs a human/AI loop
func (g *game) updateGame() {
	input(g)
	if g.gui.newGame == false && g.won == false && g.gui.drawIntro == true {
		if isPlayerHuman(g) == true || isPlayerHotseat(g) == true {
			humanLoop(g)
		} else {
			aiLoop(g)
		}
	}
}

// Play initializes a new game and launches the GUI (Ebiten)
func Play() {
	g := newGame()
	g.gui.newGame = true
	runGui()
}
