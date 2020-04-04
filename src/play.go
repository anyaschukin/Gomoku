package gomoku

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// guiReset prepares gui for the upcoming move
func guiReset(g *game) {
	g.gui.capturedPositions2 = g.gui.capturedPositions
	g.gui.capturedPositions = nil
	if alignFive(g.winMove, &g.goban, &g.align5, opponent(g.player), g.capture0, g.capture1) == false {
		g.gui.align5Positions = nil
	}
	g.gui.canCapturePositions = nil
	g.gui.canCaptureByPlaying = nil
}

// swapPlayers prepares for the next move
func swapPlayers(coordinate coordinate, g *game) {
	swapBool(&g.player)
	g.lastMove2 = g.lastMove
	g.lastMove = coordinate
	if g.won == false {
		g.gui.message = ""
	}
	if capturedTen(g) == true {
		g.gui.align5Positions = nil
	}
	g.move++
	if isPlayerHotseat(g) == true {
		aiSuggestMove(g)
	}
	captureCheat(&g.goban, g.player)
}

// gameLoop runs one move
func gameLoop(coordinate coordinate, g *game) {
	guiReset(g)
	validated := placeIfValid(coordinate, g)
	if validated == true {
		capture(coordinate, g)
		checkWin(coordinate, g)
		swapPlayers(coordinate, g)
	}
}

// humanMove listens for a click on the goban runs gameloop with clicked coordinate
func humanMove(g *game) {
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

// aiMove listens runs gameloop with suggested coordinate
func aiMove(g *game) {
	aiSuggestMove(g)
	coordinate := g.ai0.suggest
	if g.player == true {
		coordinate = g.ai1.suggest
	}
	gameLoop(coordinate, g)
}

// updateGame listens for input, and runs a human/AI move
func (g *game) updateGame() {
	input(g)
	if g.gui.newGame == false && g.won == false && g.gui.drawIntro == true {
		if isPlayerHuman(g) == true || isPlayerHotseat(g) == true {
			humanMove(g)
		} else {
			aiMove(g)
		}
	}
}
