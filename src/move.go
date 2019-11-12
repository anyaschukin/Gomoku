package play

func placeStone(coordinate coordinate, player bool, Goban *[19][19]position) {
	Goban[coordinate.y][coordinate.x].occupied = true
	Goban[coordinate.y][coordinate.x].player = player
}

func isMoveValid(coordinate coordinate, g *game) bool {
	if positionOccupied(coordinate, &g.Goban) == true {
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
		placeStone(coordinate, g.player, &g.Goban)
		return true
	}
	return false
}
