package play

func placeStone(coordinate coordinate, player bool, goban *[19][19]position) {
	goban[coordinate.y][coordinate.x].occupied = true
	goban[coordinate.y][coordinate.x].player = player
}

func isMoveValid(coordinate coordinate, g *game) bool {
	if positionOccupied(coordinate, &g.goban) == true {
		g.gui.message = "Position Occupied!"
		return false
	}
	if doubleThree(coordinate, g) == true {
		g.gui.message = "Double-Three!"
		return false
	}
	return true
}

func placeIfValid(coordinate coordinate, g *game) bool {
	if isMoveValid(coordinate, g) == true {
		placeStone(coordinate, g.player, &g.goban)
		return true
	}
	return false
}

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
