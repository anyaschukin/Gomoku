package play

// capturedEight returns true if given player has already captured 8
func capturedEight(player bool, capture0 uint8, capture1 uint8) bool {
	if player == false {
		if capture0 >= 8 {
			return true
		}
	} else {
		if capture1 >= 8 {
			return true
		}
	}
	return false
}

// cancaptureVertex returns true if given coordinate can capture on given vertex in the next move
func cancaptureVertex(coordinate coordinate, goban *[19][19]position, y int8, x int8, player bool) bool {
	one := findNeighbour(coordinate, y, x, 1)
	two := findNeighbour(coordinate, y, x, 2)
	three := findNeighbour(coordinate, y, x, 3)
	if positionOccupiedByOpponent(one, goban, player) == true &&
		positionOccupiedByOpponent(two, goban, player) == true &&
		positionUnoccupied(three, goban) == true {
		// fmt.Printf("capture possible! Player: %v can capture y:%d x:%d & y:%d x:%d\n\n", Player, one.y, one.x, two.y, two.x) /// tips flag!!
		return true
	}
	return false
}

// cancapture returns true if given coordinate can capture in the next move
func cancapture(coordinate coordinate, goban *[19][19]position, player bool) bool {
	var y int8
	var x int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				if cancaptureVertex(coordinate, goban, y, x, player) == true {
					return true
				}
			}
		}
	}
	return false
}

// captureAvailable returns true if given Player can capture in the next move (iterate entire goban, check if capture possible for each positon)
func captureAvailable(goban *[19][19]position, player bool) bool {
	var y int8
	var x int8
	for y = 0; y < 19; y++ {
		for x = 0; x < 19; x++ {
			coordinate := coordinate{y, x}
			if positionOccupiedByPlayer(coordinate, goban, player) == true {
				if cancapture(coordinate, goban, player) == true {
					return true
				}
			}
		}
	}
	return false
}

// canWinBycapture returns true if is it possible for the opponent to win by capturing 10. (have they already captured 8, and is there an available capture move)
func canWinBycapture(goban *[19][19]position, player bool, capture0 uint8, capture1 uint8) bool {
	if capturedEight(player, capture0, capture1) == true &&
		captureAvailable(goban, player) == true {
		return true
	}
	return false
}
