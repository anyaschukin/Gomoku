package play

func placeStone(coordinate coordinate, player bool, goban *[19][19]position) {
	goban[coordinate.y][coordinate.x].occupied = true
	goban[coordinate.y][coordinate.x].player = player
}

func isMoveValid(coordinate coordinate, g *game) bool {
	if positionOccupied(coordinate, &g.goban) == true {
		g.Message = "Position Occupied"
		return false
	}
	if doubleThree(coordinate, g) == true {
		g.Message = "Double-Three"
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
