package play

func coordinateOnGoban(coordinate coordinate) bool {
	if coordinate.y < 0 || coordinate.y > 18 ||
		coordinate.x < 0 || coordinate.x > 18 {
		return false
	}
	return true
}

func positionOccupied(coordinate coordinate, goban *[19][19]position) bool {
	if goban[coordinate.y][coordinate.x].occupied == true {
		return true
	}
	return false
}

func samePlayer(coordinate coordinate, goban *[19][19]position, player bool) bool {
	if goban[coordinate.y][coordinate.x].player == player {
		return true
	}
	return false
}

func positionOccupiedByPlayer(coordinate coordinate, goban *[19][19]position, player bool) bool {
	if coordinateOnGoban(coordinate) == true {
		if positionOccupied(coordinate, goban) == true &&
			samePlayer(coordinate, goban, player) == true {
			return true
		}
	}
	return false
}

func positionOccupiedByOpponent(coordinate coordinate, goban *[19][19]position, player bool) bool {
	if coordinateOnGoban(coordinate) == true {
		if positionOccupied(coordinate, goban) == true &&
			samePlayer(coordinate, goban, player) == false {
			return true
		}
	}
	return false
}

func positionUnoccupied(coordinate coordinate, goban *[19][19]position) bool {
	if coordinateOnGoban(coordinate) == true {
		if positionOccupied(coordinate, goban) == false {
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
