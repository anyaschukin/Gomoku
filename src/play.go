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
	g.DrawLastMove = true /////////// implement in gui!!!!!!!
	suggestMove(g)
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
	if g.Won == false {
		g.Message = ""
	}
	g.move++

}

func isPlayerHuman(g *game) bool {
	if (g.player == false && g.ai0.aiPlayer == false) ||
		(g.player == true && g.ai1.aiPlayer == false) {
		return true
	}
	return false
}

func isPlayerHotseat(g *game) bool {
	if (g.player == false && g.ai0.hotseat == true) ||
		(g.player == true && g.ai1.hotseat == true) {
		return true
	}
	return false
}

// suggestMove checks if Player is ai & if so prompts the ai to suggest a move
func suggestMove(g *game) {
	if isPlayerHuman(g) == false || isPlayerHotseat(g) == true {
		artificialIdiot(g) /////create move suggestion
	}
}

// gameLoop runs one move
func gameLoop(coordinate coordinate, g *game) {
	validated := placeIfValid(coordinate, g)
	if validated == true {
		capture(coordinate, g)
		checkWin(coordinate, g)
		g.LastMove = coordinate
		swapPlayers(g)
	}
	suggestMove(g)
}

// humanLoop listens for a click on the goban runs gameloop with clicked coordinate
func humanLoop(g *game) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) == true {
		x, y := ebiten.CursorPosition()
		if clickgoban(x, y) == true {
			coordinate := coordinate{-1, -1}
			coordinate.x = int8((float64(x) - (zerox * scale)) / (positionWidth * scale))
			coordinate.y = int8((float64(y) - (zeroy * scale)) / (positionWidth * scale))
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

func (g *game) updateGame() { ////listen for input, update struct
	input(g)
	if g.newGame == false && g.Won == false {
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
	runEbiten()
}
