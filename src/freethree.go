package play

func threeBlocked(end1 coordinate, end2 coordinate, Goban *[19][19]position) bool {
	if positionUnoccupied(end1, Goban) == true &&
		positionUnoccupied(end2, Goban) == true {
		return false
	}
	return true
}

// checkVertexForThree returns true if it finds an unblocked FreeThree on given vertex
func checkVertexForThree(coordinate coordinate, Goban *[19][19]position, y int8, x int8, player bool) bool {
	minusTwo := findNeighbour(coordinate, y, x, -2)
	minusOne := findNeighbour(coordinate, y, x, -1)
	one := findNeighbour(coordinate, y, x, 1)
	two := findNeighbour(coordinate, y, x, 2)
	three := findNeighbour(coordinate, y, x, 3)
	four := findNeighbour(coordinate, y, x, 4)
	if positionOccupiedByPlayer(one, Goban, player) == true {
		if positionOccupiedByPlayer(two, Goban, player) == true {
			if threeBlocked(minusOne, three, Goban) == false {
				return true
			}
		}
		if positionOccupiedByPlayer(three, Goban, player) == true {
			if threeBlocked(minusOne, four, Goban) == false {
				if positionOccupiedByOpponent(two, Goban, player) == false {
					return true
				}
			}
		}
		if y < 0 || (y == 0 && x == -1) {
			if positionOccupiedByPlayer(minusOne, Goban, player) == true {
				if threeBlocked(minusTwo, two, Goban) == false {
					return true
				}
			}
		}
	}
	if positionOccupiedByPlayer(two, Goban, player) == true {
		if positionOccupiedByPlayer(three, Goban, player) == true {
			if threeBlocked(minusOne, four, Goban) == false {
				if positionOccupiedByOpponent(one, Goban, player) == false {
					return true
				}
			}
		}
		if positionOccupiedByPlayer(minusOne, Goban, player) == true {
			if threeBlocked(minusTwo, three, Goban) == false {
				if positionOccupiedByOpponent(one, Goban, player) == false {
					return true
				}
			}
		}
	}
	return false
}

// doubleThree returns false if suggested move breaks the double three rule
func doubleThree(coordinate coordinate, g *game) (valid bool) {
	var freeThree bool
	var x int8
	var y int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				foundThree := checkVertexForThree(coordinate, &g.Goban, y, x, g.player)
				if foundThree == true {
					if freeThree == true {
						return true
					} else {
						freeThree = true
					}
				}
			}
		}
	}
	return false
}
