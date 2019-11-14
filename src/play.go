package play

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// newGame initializes a new game
func newGame() *game {
	g = &game{}
	g.ai0.aiPlayer = true
	g.ai0.depth = 3
	g.ai1.depth = 3
	g.drawLastMove = true //// false for correction (bonus)!!!!
	g.drawWinMove = true  //// false for correction (bonus)!!!!
	aiSuggestMove(g)
	return g
}

// opponent returns the opponent of the current Player
func opponent(player bool) bool {
	if player == false {
		return true
	}
	return false
}

// swapPlayers swaps Players, clears the message and iterates move
func swapPlayers(g *game) {
	g.player = opponent(g.player)
	if g.won == false {
		g.message = ""
	}
	g.move++

}

// gameLoop runs one move
func gameLoop(coordinate coordinate, g *game) {
	validated := placeIfValid(coordinate, g)
	if validated == true {
		capture(coordinate, g)
		checkWin(coordinate, g)
		g.lastMove = coordinate
		swapPlayers(g)
	}
	aiSuggestMove(g)
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
}

// updateGame listens for input, and runs a human/AI loop
func (g *game) updateGame() {
	input(g)
	if g.newGame == false && g.won == false {
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
	g.ai1.depth = 3 ////////
	runGui()
}
