package play

func coordinateOnGoban(coordinate coordinate) (onGoban bool) {
	if coordinate.y < 0 || coordinate.y > 18 || coordinate.x < 0 || coordinate.x > 18 {
		return false
	}
	return true
}

func positionOccupied(coordinate coordinate, Goban *[19][19]position) (occupied bool) {
	if Goban[coordinate.y][coordinate.x].occupied == true {
		return true
	}
	return false
}

func samePlayer(coordinate coordinate, Goban *[19][19]position, player bool) (samePlayer bool) {
	if Goban[coordinate.y][coordinate.x].player == player {
		return true
	}
	return false
}

func positionOccupiedByPlayer(coordinate coordinate, Goban *[19][19]position, player bool) bool {
	if coordinateOnGoban(coordinate) == true {
		if positionOccupied(coordinate, Goban) == true &&
			samePlayer(coordinate, Goban, player) == true {
			return true
		}
	}
	return false
}

func positionOccupiedByOpponent(coordinate coordinate, Goban *[19][19]position, player bool) bool {
	if coordinateOnGoban(coordinate) == true {
		if positionOccupied(coordinate, Goban) == true &&
			samePlayer(coordinate, Goban, player) == false {
			return true
		}
	}
	return false
}

func positionUnoccupied(coordinate coordinate, Goban *[19][19]position) (unoccupied bool) {
	if coordinateOnGoban(coordinate) == true {
		if positionOccupied(coordinate, Goban) == false {
			return true
		}
	}
	return false
}

func findNeighbour(coordinate coordinate, y int8, x int8, multiple int8) coordinate {
	neighbour := coordinate
	neighbour.y += y * multiple
	neighbour.x += x * multiple
	return neighbour
}
