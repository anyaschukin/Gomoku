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
		if PositionOccupied(coordinate, goban) == true {
			if SamePlayer(coordinate, goban, player) == true {
				return true
			}
		}
	}
	return false
}

func PositionOccupiedByOpponent(coordinate coordinate, goban *[19][19]position, player bool) bool {
	if CoordinateOnGoban(coordinate) == true {
		if PositionOccupied(coordinate, goban) == true {
			if SamePlayer(coordinate, goban, player) == false {
				return true
			}
		}
	}
	return false
}

func threeBlocked(end1 coordinate, end2 coordinate, goban *[19][19]position, player bool) bool {
	if PositionOccupiedByOpponent(end1, goban, player) == false && PositionOccupiedByOpponent(end2, goban, player) == false {
		return false
	}
	return true
}

func FindNeighbour(coordinate coordinate, y int8, x int8, multiple int8) coordinate {
	neighbour := coordinate
	neighbour.y += y * multiple
	neighbour.x += x * multiple
	return neighbour
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
