package play

// ultimately, I'd like to re-use the same functions that Drew wrote originally
// that way, we're DRY

// re-write these functions so you're passing the goban and player, but not the whole game!

// doubleThree returns true if suggested move breaks the double three rule
func doubleThree2(coordinate coordinate, goban *[19][19]position, player bool) bool {
	var freeThree bool
	var y int8
	var x int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				foundThree := checkVertexForThree(coordinate, goban, y, x, player)
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


func isMoveValid2(coordinate coordinate, goban *[19][19]position, player bool) bool {
	if positionOccupied(coordinate, goban) == true {
		// g.message = "Position Occupied"
		return false
	}
	if doubleThree2(coordinate, goban, player) == true {	 // duplicate w/o *game
		// g.message = "Double-Three"
		return false
	}
	return true
}
