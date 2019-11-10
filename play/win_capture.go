package play

// capturedEight returns true if given Player has already captured 8
func capturedEight(Player bool, capture0 uint8, capture1 uint8) bool {
	if Player == false {
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

// canCaptureVertex returns true if given coordinate can capture on given vertex in the next move
func canCaptureVertex(coordinate coordinate, Goban *[19][19]position, y int8, x int8, Player bool) bool {
	one := FindNeighbour(coordinate, y, x, 1)
	two := FindNeighbour(coordinate, y, x, 2)
	three := FindNeighbour(coordinate, y, x, 3)
	if PositionOccupiedByOpponent(one, Goban, Player) == true &&
		PositionOccupiedByOpponent(two, Goban, Player) == true &&
		PositionUnoccupied(three, Goban) == true {
		// fmt.Printf("Capture possible! Player: %v can capture y:%d x:%d & y:%d x:%d\n\n", Player, one.y, one.x, two.y, two.x) /// tips flag!!
		return true
	}
	return false
}

// canCapture returns true if given coordinate can capture in the next move
func canCapture(coordinate coordinate, Goban *[19][19]position, Player bool) bool {
	var y int8
	var x int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				if canCaptureVertex(coordinate, Goban, y, x, Player) == true {
					return true
				}
			}
		}
	}
	return false
}

// captureAvailable returns true if given Player can capture in the next move (iterate entire Goban, check if capture possible for each positon)
func captureAvailable(Goban *[19][19]position, Player bool) bool {
	var y int8
	var x int8
	for y = 0; y < 19; y++ {
		for x = 0; x < 19; x++ {
			coordinate := coordinate{y, x}
			if PositionOccupiedByPlayer(coordinate, Goban, Player) == true {
				if canCapture(coordinate, Goban, Player) == true {
					return true
				}
			}
		}
	}
	return false
}

// canWinByCapture returns true if is it possible for the opponent to win by capturing 10. (have they already captured 8, and is there an available capture move)
func canWinByCapture(Goban *[19][19]position, Player bool, capture0 uint8, capture1 uint8) bool {
	if capturedEight(Player, capture0, capture1) == true &&
		captureAvailable(Goban, Player) == true {
		return true
	}
	return false
}
