package gomoku

// willCaptureVertex returns true if given coordinate will capture in the next move
func willCapture(coordinate coordinate, goban *[19][19]position, y, x, i int8, player bool) bool {
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

// willCaptureDirection returns true if given coordinate will capture in given direction in the next move
func willCaptureDirection(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
	if willCapture(coordinate, goban, y, x, 1, player) == true || willCapture(coordinate, goban, y, x, -1, player) == true {
		return true
	}
	return false
}

// chainLength returns the total length of stones aligned running through given a coordinate on a given axe
func chainLength(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) int8 {
	// 	// if positionOccupiedByPlayer(coordinate, &g.goban, player) == false {
	// 	// 	return false
	a := measureChain(coordinate, goban, y, x, player)
	b := measureChain(coordinate, goban, -y, -x, player)
	return a + b + 1
}
