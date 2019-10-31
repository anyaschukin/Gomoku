package play

// canBeCapturedVertex returns true if given coordinate can be captured on given vertex in the next move
func canBeCapturedVertex(coordinate coordinate, goban *[19][19]position, y int8, x int8, player bool) bool {
	minusOne := FindNeighbour(coordinate, y, x, -1)
	one := FindNeighbour(coordinate, y, x, 1)
	two := FindNeighbour(coordinate, y, x, 2)
	if PositionOccupiedByPlayer(one, goban, player) {
		if PositionOccupiedByOpponent(minusOne, goban, player) && PositionUnoccupied(two, goban) {
			return true
		}
		if PositionOccupiedByOpponent(two, goban, player) && PositionUnoccupied(minusOne, goban) {
			return true
		}
	}
	return false
}

// canBeCapturedVertices returns true if given coordinate can be captured in the next move
func canBeCapturedVertices(coordinate coordinate, goban *[19][19]position, player bool) bool {
	var y int8
	var x int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				if canBeCapturedVertex(coordinate, goban, y, x, player) == true {
					return true
				}
			}
		}
	}
	return false
}

// CanBreakFive returns true if its possible to break the aligned 5
func canBreakFive(coordinate coordinate, goban *[19][19]position, y int8, x int8, player bool) bool {
	if canBeCapturedVertices(coordinate, goban, player) == true {
		return true
	}
	//move along winning string//////////////
	var multiple int8
	var a int8
	var b int8
	for multiple = 1; multiple < 5; multiple++ {
		neighbour := FindNeighbour(coordinate, y, x, multiple)
		if PositionOccupiedByPlayer(neighbour, goban, player) == true &&
			canBeCapturedVertices(neighbour, goban, player) == false {
			a++
		} else {
			break
		}
	}
	for multiple = -1; multiple > -5; multiple-- {
		neighbour := FindNeighbour(coordinate, y, x, multiple)
		if PositionOccupiedByPlayer(neighbour, goban, player) == true &&
			canBeCapturedVertices(neighbour, goban, player) == false {
			b++
		} else {
			break
		}
	}
	if a+b >= 4 {
		return false
	}
	return true
}
