package gomoku

// canBeCapturedVertex returns true if given coordinate can be captured on given vertex in the next move
func canBeCapturedVertex(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
	minusOne := findNeighbour(coordinate, y, x, -1)
	one := findNeighbour(coordinate, y, x, 1)
	two := findNeighbour(coordinate, y, x, 2)
	if positionOccupiedByPlayer(one, goban, player) {
		if (positionOccupiedByOpponent(minusOne, goban, player) && positionUnoccupied(two, goban)) ||
			(positionOccupiedByOpponent(two, goban, player) && positionUnoccupied(minusOne, goban)) {
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

// measureChainBeCaptured returns how many stones in a row cannot be captured
func measureChainBeCaptured(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) int8 {
	var length int8
	var multiple int8
	for multiple = 1; multiple < 5; multiple++ {
		neighbour := findNeighbour(coordinate, y, x, multiple)
		if positionOccupiedByPlayer(neighbour, goban, player) == true &&
			canBeCapturedVertices(neighbour, goban, player) == false {
			length++
		} else {
			break
		}
	}
	return length
}

// canBreakFive returns true if its possible to break the aligned 5
func canBreakFive(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
	if canBeCapturedVertices(coordinate, goban, player) == true {
		return true
	}
	a := measureChainBeCaptured(coordinate, goban, y, x, player)
	b := measureChainBeCaptured(coordinate, goban, -y, -x, player)
	if a+b >= 4 {
		return false
	}
	return true
}
