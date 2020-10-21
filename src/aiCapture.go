package gomoku

// breakFive returns true if placing player's stone at coordinate will break opponent's five-in-a-row
func breakFive(coordinate coordinate, goban *[19][19]position, player bool) bool {
	var x int8
	var y int8
	for y = -1; y <= 0; y++ {
		for x = -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				return false
			}
			if checkVertexAlignFive(coordinate, goban, y, x, player) == true {
				return true
			}
		}
	}
	return false
}

// breakFiveDirection returns true if placing player's stone at coordinate will break opponent's five-in-a-row
func breakFiveDirection(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
	one := findNeighbour(coordinate, y, x, 1)
	two := findNeighbour(coordinate, y, x, 2)
	if (positionOccupiedByOpponent(one, goban, player) == true && breakFive(one, goban, !player) == true) ||
		(positionOccupiedByOpponent(two, goban, player) == true && breakFive(two, goban, !player) == true) {
		return true
	}
	return false
}

//  willBreak5Align returns true if placing player's stone at coordinate will break opponent's five-in-a-row
func willBreak5Align(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
	if breakFiveDirection(coordinate, goban, y, x, player) == true || breakFiveDirection(coordinate, goban, -y, -x, player) == true {
		// fmt.Printf("willBreak5Align: coordinate = %v, player = %v\n", coordinate, player)
		return true
	}
	return false
}

// willCaptureVertex returns true if given coordinate will capture in the next move
func willCaptureDirection(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
	one := findNeighbour(coordinate, y, x, 1)
	two := findNeighbour(coordinate, y, x, 2)
	three := findNeighbour(coordinate, y, x, 3)
	if positionOccupiedByOpponent(one, goban, player) == true &&
		positionOccupiedByOpponent(two, goban, player) == true &&
		positionOccupiedByPlayer(three, goban, player) == true {
		return true
	}
	return false
}

// willCapture returns number of captures given coordinate will capture (for player) in given vertex in the next move
func willCapture(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) uint8 {
	var cap uint8
	if willCaptureDirection(coordinate, goban, y, x, player) == true {
		cap++
	}
	if willCaptureDirection(coordinate, goban, -y, -x, player) == true {
		cap++
	}
	return cap
}

// willBeCaptured returns true if given coordinate will be captured (for opponent) in given direction in the next move
func willBeCaptured(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
	if canBeCapturedVertex(coordinate, goban, y, x, player) == true || canBeCapturedVertex(coordinate, goban, -y, -x, player) == true {
		return true
	}
	return false
}

// captureAttackDefend returns a score for capturing 2 stones, or for protecting against being captured
func captureAttackDefend(coordinate coordinate, goban *[19][19]position, y, x int8, player bool, captures captures) int {
	cap := willCapture(coordinate, goban, y, x, player)
	if cap != 0 {
		if capturedEight(player, captures.capture0, captures.capture1) == true { // this is the ultimate winning move
			return capture10
		} else if willBreak5Align(coordinate, goban, y, x, player) == true {
			return break5Align
		}
		if cap == 2 {
			return capture2 * 2
		}
		return capture2
	} else if willBeCaptured(coordinate, goban, y, x, player) == true {
		if capturedEight(!player, captures.capture0, captures.capture1) == true {
			return willBeCaptured8
		}
		return willBeCaptured2
	}
	return 0
}
