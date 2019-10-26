package play

func threeBlocked(end1 coordinate, end2 coordinate, goban *[19][19]position, player bool) bool {
	if PositionOccupiedByOpponent(end1, goban, player) == false && PositionOccupiedByOpponent(end2, goban, player) == false {
		return false
	}
	return true
}

func checkVertexForThree(coordinate coordinate, goban *[19][19]position, y int8, x int8, player bool) (FreeThree bool) {
	minusTwo := FindNeighbour(coordinate, y, x, -2)
	minusOne := FindNeighbour(coordinate, y, x, -1)
	one := FindNeighbour(coordinate, y, x, 1)
	two := FindNeighbour(coordinate, y, x, 2)
	three := FindNeighbour(coordinate, y, x, 3)
	four := FindNeighbour(coordinate, y, x, 4)
	if PositionOccupiedByPlayer(one, goban, player) == true {
		if PositionOccupiedByPlayer(two, goban, player) == true {
			if threeBlocked(minusOne, three, goban, player) == false {
				return true
			}
		}
		if PositionOccupiedByPlayer(three, goban, player) == true {
			if threeBlocked(minusOne, four, goban, player) == false {
				if PositionOccupiedByOpponent(two, goban, player) == false {
					return true
				}
			}
		}
		if y < 0 || (y == 0 && x == -1) {
			if PositionOccupiedByPlayer(minusOne, goban, player) == true {
				if threeBlocked(minusTwo, two, goban, player) == false {
					return true
				}
			}
		}
	}
	if PositionOccupiedByPlayer(two, goban, player) == true {
		if PositionOccupiedByPlayer(three, goban, player) == true {
			if threeBlocked(minusOne, four, goban, player) == false {
				if PositionOccupiedByOpponent(one, goban, player) == false {
					return true
				}
			}
		}
		if PositionOccupiedByPlayer(minusOne, goban, player) == true {
			if threeBlocked(minusTwo, three, goban, player) == false {
				if PositionOccupiedByOpponent(one, goban, player) == false {
					return true
				}
			}
		}
	}
	return false
}

func DoubleThree(coordinate coordinate, g *game) (valid bool) { // returns false if move breaks rule
	var freeThree bool
	var x int8
	var y int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				foundThree := checkVertexForThree(coordinate, &g.goban, y, x, g.player)
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
