package play

import (
	"github.com/hajimehoshi/ebiten"
)

// undo undoes the last 2 moves
func undo(g *game) {
	if g.gui.undo == true && g.won == false &&
		g.move > 1 && g.move > g.gui.undoMove {
		g.gui.undoMove = g.move
		for _, position := range g.gui.capturedPositions {
			placeStone(position, g.player, &g.goban)
		}
		for _, position := range g.gui.capturedPositions2 {
			placeStone(position, opponent(g.player), &g.goban)
		}
		removeStone(g.lastMove, &g.goban)
		removeStone(g.lastMove2, &g.goban)
		g.move -= 2
	}
}

// recordAlignFive stores the positions of the align 5 for display in gui
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

func drawAlign5(screen *ebiten.Image, g *game, alpha float64) {
	if g.won == true || g.gui.tips == true {
		for _, position := range g.gui.align5Positions {
			drawImagePulse(screen, imgRed, stoneX(position.x), stoneY(position.y), scale, alpha)
		}
	}
}
