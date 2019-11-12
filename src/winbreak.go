package play

// canBecapturedVertex returns true if given coordinate can be captured on given vertex in the next move
func canBecapturedVertex(coordinate coordinate, Goban *[19][19]position, y int8, x int8, player bool) bool {
	minusOne := findNeighbour(coordinate, y, x, -1)
	one := findNeighbour(coordinate, y, x, 1)
	two := findNeighbour(coordinate, y, x, 2)
	if positionOccupiedByPlayer(one, Goban, player) {
		if positionOccupiedByOpponent(minusOne, Goban, player) && positionUnoccupied(two, Goban) {
			return true
		}
		if positionOccupiedByOpponent(two, Goban, player) && positionUnoccupied(minusOne, Goban) {
			return true
		}
	}
	return false
}

// canBecapturedVertices returns true if given coordinate can be captured in the next move
func canBecapturedVertices(coordinate coordinate, Goban *[19][19]position, player bool) bool {
	var y int8
	var x int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				if canBecapturedVertex(coordinate, Goban, y, x, player) == true {
					return true
				}
			}
		}
	}
	return false
}

// CanBreakFive returns true if its possible to break the aligned 5
func canBreakFive(coordinate coordinate, Goban *[19][19]position, y int8, x int8, player bool) bool {
	if canBecapturedVertices(coordinate, Goban, player) == true {
		return true
	}
	//move along winning string//////////////
	var multiple int8
	var a int8
	var b int8
	for multiple = 1; multiple < 5; multiple++ {
		neighbour := findNeighbour(coordinate, y, x, multiple)
		if positionOccupiedByPlayer(neighbour, Goban, player) == true &&
			canBecapturedVertices(neighbour, Goban, player) == false {
			a++
		} else {
			break
		}
	}
	for multiple = -1; multiple > -5; multiple-- {
		neighbour := findNeighbour(coordinate, y, x, multiple)
		if positionOccupiedByPlayer(neighbour, Goban, player) == true &&
			canBecapturedVertices(neighbour, Goban, player) == false {
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
