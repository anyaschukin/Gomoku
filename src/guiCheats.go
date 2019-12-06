package gomoku

import (
	"github.com/hajimehoshi/ebiten"
)

// undo undoes the last 2 moves
func undo(g *game) {
	if g.gui.undo == true && (isPlayerHuman(g) || isOpponentHuman(g)) &&
		g.move > 1 && g.move >= g.gui.undoMove {
		g.gui.undoMove = g.move
		for _, position := range g.gui.capturedPositions {
			placeStone(position, g.player, &g.goban)
		}
		for _, position := range g.gui.capturedPositions2 {
			placeStone(position, opponent(g.player), &g.goban)
		}
		removeStone(g.lastMove, &g.goban)
		removeStone(g.lastMove2, &g.goban)
		g.gui.capturedPositions = nil
		g.gui.capturedPositions2 = nil
		g.move -= 2
		g.won = false
		g.gui.message = ""
	}
}

// recordAlignFive stores align 5 positions for display in gui
func recordAlignFive(coordinate coordinate, y, x int8, g *game) {
	g.winMove = coordinate
	g.gui.align5Positions = append(g.gui.align5Positions, coordinate)
	var multiple int8
	var a int8
	var b int8
	for multiple = 1; multiple < 5; multiple++ {
		neighbour := findNeighbour(coordinate, y, x, multiple)
		if positionOccupiedByPlayer(neighbour, &g.goban, g.player) == true {
			g.gui.align5Positions = append(g.gui.align5Positions, neighbour)
			a++
		} else {
			break
		}
	}
	for multiple = -1; multiple > -5; multiple-- {
		neighbour := findNeighbour(coordinate, y, x, multiple)
		if positionOccupiedByPlayer(neighbour, &g.goban, g.player) == true {
			g.gui.align5Positions = append(g.gui.align5Positions, neighbour)
			b++
		} else {
			break
		}
	}
}

// drawAlign5 displays pulsing red align5 for tips
func drawAlign5(screen *ebiten.Image, g *game, alpha float64) {
	if g.won == true || g.gui.tips == true {
		for _, position := range g.gui.align5Positions {
			drawImagePulse(screen, imgRed, stoneX(position.x), stoneY(position.y), scale, alpha)
		}
	}
}

// canCaptureCheat stores positions if a capture is possible for a given position and vertex
func canCaptureVertexCheat(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) {
	one := findNeighbour(coordinate, y, x, 1)
	two := findNeighbour(coordinate, y, x, 2)
	three := findNeighbour(coordinate, y, x, 3)
	if positionOccupiedByOpponent(one, goban, player) == true &&
		positionOccupiedByOpponent(two, goban, player) == true &&
		positionUnoccupied(three, goban) == true {
		g.gui.canCapturePositions = append(g.gui.canCapturePositions, one, two)
		g.gui.canCaptureByPlaying = append(g.gui.canCaptureByPlaying, three)
	}
}

// canCaptureCheat checks if a capture is possible for a given position
func canCaptureCheat(coordinate coordinate, goban *[19][19]position, player bool) {
	var y int8
	var x int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				canCaptureVertexCheat(coordinate, goban, y, x, player)
			}
		}
	}
}

// captureCheat stores possible capture positions for gui tips
// (iterate entire goban, check if capture possible for each positon)
func captureCheat(goban *[19][19]position, player bool) {
	if g.gui.tips == true {
		var y int8
		var x int8
		for y = 0; y < 19; y++ {
			for x = 0; x < 19; x++ {
				coordinate := coordinate{y, x}
				if positionOccupiedByPlayer(coordinate, goban, player) == true {
					canCaptureCheat(coordinate, goban, player)
				}
			}
		}
	}
}

// drawCaptureCheat draws pulsing possible captures for gui tips
func drawCaptureCheat(screen *ebiten.Image, g *game, alpha float64) {
	if g.gui.tips == true && isPlayerHuman(g) == true && g.won == false {
		for _, position := range g.gui.canCapturePositions {
			drawCapturedPosition(screen, g, position, alpha)
		}
		for _, position := range g.gui.canCaptureByPlaying {
			if g.player == false {
				drawImagePulse(screen, imgBlack, stoneX(position.x), stoneY(position.y), scale, alpha)
			} else {
				drawImagePulse(screen, imgWhite, stoneX(position.x), stoneY(position.y), scale, alpha)
			}
		}
	}
}

// drawTips draws tips in gui
func drawTips(screen *ebiten.Image, g *game, alpha float64) {
	drawAlign5(screen, g, alpha)
	drawCaptureCheat(screen, g, alpha)
}
