package play

// canBeCapturedVertex returns true if given coordinate can be captured on given vertex in the next move
func canBeCapturedVertex(coordinate coordinate, Goban *[19][19]position, y int8, x int8, Player bool) bool {
	minusOne := FindNeighbour(coordinate, y, x, -1)
	one := FindNeighbour(coordinate, y, x, 1)
	two := FindNeighbour(coordinate, y, x, 2)
	if PositionOccupiedByPlayer(one, Goban, Player) {
		if PositionOccupiedByOpponent(minusOne, Goban, Player) && PositionUnoccupied(two, Goban) {
			return true
		}
		if PositionOccupiedByOpponent(two, Goban, Player) && PositionUnoccupied(minusOne, Goban) {
			return true
		}
	}
	return false
}

// canBeCapturedVertices returns true if given coordinate can be captured in the next move
func canBeCapturedVertices(coordinate coordinate, Goban *[19][19]position, Player bool) bool {
	var y int8
	var x int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				if canBeCapturedVertex(coordinate, Goban, y, x, Player) == true {
					return true
				}
			}
		}
	}
	return false
}

// CanBreakFive returns true if its possible to break the aligned 5
func canBreakFive(coordinate coordinate, Goban *[19][19]position, y int8, x int8, Player bool) bool {
	if canBeCapturedVertices(coordinate, Goban, Player) == true {
		return true
	}
	//move along winning string//////////////
	var multiple int8
	var a int8
	var b int8
	for multiple = 1; multiple < 5; multiple++ {
		neighbour := FindNeighbour(coordinate, y, x, multiple)
		if PositionOccupiedByPlayer(neighbour, Goban, Player) == true &&
			canBeCapturedVertices(neighbour, Goban, Player) == false {
			a++
		} else {
			break
		}
	}
	for multiple = -1; multiple > -5; multiple-- {
		neighbour := FindNeighbour(coordinate, y, x, multiple)
		if PositionOccupiedByPlayer(neighbour, Goban, Player) == true &&
			canBeCapturedVertices(neighbour, Goban, Player) == false {
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
