package gomoku

func isMoveValid2(coordinate coordinate, goban *[19][19]position, player bool) bool {
	if coordinateOnGoban(coordinate) == false {
		return false
	}
	if positionOccupied(coordinate, goban) == true {
		return false
	}
	if doubleThree(coordinate, goban, !player) == true {
		return false
	}
	return true
}

// Returns true if given position has immediate neighbor which is occupied
// Optimization so that only populated parts of the board are explored. Standalone/isolated positions are ignored.
func hasNeigbours(y_orig int8, x_orig int8, goban *[19][19]position) bool {
	possibleMove := coordinate{y_orig, x_orig}
	if coordinateOnGoban(possibleMove) == false {
		return false
	}
	var x int8
	var y int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				neighbour := findNeighbour(possibleMove, y, x, 1)
				if coordinateOnGoban(neighbour) == true {
					if positionOccupied(neighbour, goban) == true {
						return true
					}
				}
			}
		}
	}
	return false
}
