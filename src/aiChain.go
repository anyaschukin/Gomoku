package gomoku

// willCaptureVertex returns true if given coordinate will capture in the next move
func willCaptureDirection(coordinate coordinate, goban *[19][19]position, y, x, i int8, player bool) bool {
	one := findNeighbour(coordinate, y, x, i*1)
	two := findNeighbour(coordinate, y, x, i*2)
	three := findNeighbour(coordinate, y, x, i*3)
	if positionOccupiedByOpponent(one, goban, player) == true &&
		positionOccupiedByOpponent(two, goban, player) == true &&
		positionOccupiedByPlayer(three, goban, player) == true {
		return true
	}
	return false
}

// willCapture returns true if given coordinate will capture (for player) in given direction in the next move
func willCapture(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
	if willCaptureDirection(coordinate, goban, y, x, 1, player) == true || willCaptureDirection(coordinate, goban, y, x, -1, player) == true {
		return true
	}
	return false
}

// willBeCaptured returns true if given coordinate will capture (for opponent) in given direction in the next move
func willBeCaptured(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
	if willCaptureDirection(coordinate, goban, y, x, 1, !player) == true || willCaptureDirection(coordinate, goban, y, x, -1, !player) == true {
		return true
	}
	return false
}

func captureOrBeCaptured(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) int {
	if willCapture(coordinate, goban, y, x, player) == true {
		// fmt.Printf("Will capture - player = %v\n", player)
		// check that canWinByCapture works

		return 42e11
	} else if canBeCapturedVertex(coordinate, goban, y, x, player) == true {
		return -42e11
	}
	return 0
}

// checks either side for whose chain it is << REWRITE THIS COMMENT
func checkNeighbors(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
	neighbour1 := findNeighbour(coordinate, y, x, -1)
	neighbour2 := findNeighbour(coordinate, y, x, 1)
	if positionOccupiedByPlayer(neighbour1, goban, player) == false || positionOccupiedByPlayer(neighbour2, goban, player) == false {
		return true
	}
	return false
}

// chainLength returns the total length of stones aligned running through given a coordinate on a given axe
func chainLength(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) int8 {
	a := measureChain(coordinate, goban, y, x, player)
	b := measureChain(coordinate, goban, -y, -x, player)
	return a + b
}

// map[string]bool
// DeepEqual?
