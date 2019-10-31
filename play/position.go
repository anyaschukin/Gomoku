package play

func CoordinateOnGoban(coordinate coordinate) (onGoban bool) {
	if coordinate.y < 0 || coordinate.y > 18 || coordinate.x < 0 || coordinate.x > 18 {
		return false
	}
	return true
}

func PositionOccupied(coordinate coordinate, goban *[19][19]position) (occupied bool) {
	if goban[coordinate.y][coordinate.x].occupied == true {
		return true
	}
	return false
}

func SamePlayer(coordinate coordinate, goban *[19][19]position, player bool) (samePlayer bool) {
	if goban[coordinate.y][coordinate.x].player == player {
		return true
	}
	return false
}

func PositionOccupiedByPlayer(coordinate coordinate, goban *[19][19]position, player bool) bool {
	if CoordinateOnGoban(coordinate) == true {
		if PositionOccupied(coordinate, goban) == true &&
			SamePlayer(coordinate, goban, player) == true {
			return true
		}
	}
	return false
}

func PositionOccupiedByOpponent(coordinate coordinate, goban *[19][19]position, player bool) bool {
	if CoordinateOnGoban(coordinate) == true {
		if PositionOccupied(coordinate, goban) == true &&
			SamePlayer(coordinate, goban, player) == false {
			return true
		}
	}
	return false
}

func PositionUnoccupied(coordinate coordinate, goban *[19][19]position) (unoccupied bool) {
	if CoordinateOnGoban(coordinate) == true {
		if PositionOccupied(coordinate, goban) == false {
			return true
		}
	}
	return false
}

func FindNeighbour(coordinate coordinate, y int8, x int8, multiple int8) coordinate {
	neighbour := coordinate
	neighbour.y += y * multiple
	neighbour.x += x * multiple
	return neighbour
}
