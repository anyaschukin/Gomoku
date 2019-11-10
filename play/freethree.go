package play

func threeBlocked(end1 coordinate, end2 coordinate, Goban *[19][19]position, Player bool) bool {
	// if PositionOccupiedByOpponent(end1, Goban, Player) == false &&
	// 	PositionOccupiedByOpponent(end2, Goban, Player) == false {
	// 	return false
	// }
	if PositionUnoccupied(end1, Goban) == true && ////////correct??
		PositionUnoccupied(end2, Goban) == true {
		return false
	}
	return true
}

// checkVertexForThree returns true if it finds an unblocked FreeThree on given vertex
func checkVertexForThree(coordinate coordinate, Goban *[19][19]position, y int8, x int8, Player bool) bool {
	minusTwo := FindNeighbour(coordinate, y, x, -2)
	minusOne := FindNeighbour(coordinate, y, x, -1)
	one := FindNeighbour(coordinate, y, x, 1)
	two := FindNeighbour(coordinate, y, x, 2)
	three := FindNeighbour(coordinate, y, x, 3)
	four := FindNeighbour(coordinate, y, x, 4)
	if PositionOccupiedByPlayer(one, Goban, Player) == true {
		if PositionOccupiedByPlayer(two, Goban, Player) == true {
			if threeBlocked(minusOne, three, Goban, Player) == false {
				return true
			}
		}
		if PositionOccupiedByPlayer(three, Goban, Player) == true {
			if threeBlocked(minusOne, four, Goban, Player) == false {
				if PositionOccupiedByOpponent(two, Goban, Player) == false {
					return true
				}
			}
		}
		if y < 0 || (y == 0 && x == -1) {
			if PositionOccupiedByPlayer(minusOne, Goban, Player) == true {
				if threeBlocked(minusTwo, two, Goban, Player) == false {
					return true
				}
			}
		}
	}
	if PositionOccupiedByPlayer(two, Goban, Player) == true {
		if PositionOccupiedByPlayer(three, Goban, Player) == true {
			if threeBlocked(minusOne, four, Goban, Player) == false {
				if PositionOccupiedByOpponent(one, Goban, Player) == false {
					return true
				}
			}
		}
		if PositionOccupiedByPlayer(minusOne, Goban, Player) == true {
			if threeBlocked(minusTwo, three, Goban, Player) == false {
				if PositionOccupiedByOpponent(one, Goban, Player) == false {
					return true
				}
			}
		}
	}
	return false
}

// DoubleThree returns false if suggested move breaks the double three rule
func DoubleThree(coordinate coordinate, g *Game) (valid bool) {
	var freeThree bool
	var x int8
	var y int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				foundThree := checkVertexForThree(coordinate, &g.Goban, y, x, g.Player)
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
