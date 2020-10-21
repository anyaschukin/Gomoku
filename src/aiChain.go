package gomoku

func measureOpponent(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) (int8, bool) {
	var multiple int8
	var length int8
	for multiple = 1; multiple < 5; multiple++ {
		neighbour := findNeighbour(coordinate, y, x, multiple)
		if positionOccupiedByOpponent(neighbour, goban, player) == true {
			length++
		} else if positionUnoccupied(neighbour, goban) == false {
			return length, true
		} else {
			break
		}
	}
	return length, false
}

// lengthOpponentChain returns the total length of player's stones aligned running through given a coordinate on a given axe
func lengthOpponentChain(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) (int8, bool, bool) {
	a, flanked1 := measureOpponent(coordinate, goban, y, x, player)
	b, flanked2 := measureOpponent(coordinate, goban, -y, -x, player)
	if a + b > 4 {
		return 4, flanked1, flanked2
	}
	return a + b, flanked1, flanked2
}

func measurePlayer(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) (int8, bool) {
	var length int8
	var multiple int8
	for multiple = 1; multiple < 5; multiple++ {
		neighbour := findNeighbour(coordinate, y, x, multiple)
		if positionOccupiedByPlayer(neighbour, goban, player) == true {
			length++
		} else if positionUnoccupied(neighbour, goban) == false {
			return length, true
		} else {
			break
		}
	}
	return length, false
}

// lengthPlayerChain returns the total length of player's stones aligned running through given a coordinate on a given axe
func lengthPlayerChain(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) (int8, bool, bool) {
	a, flanked1 := measurePlayer(coordinate, goban, y, x, player)
	b, flanked2 := measurePlayer(coordinate, goban, -y, -x, player)
	if a + b + 1 > 5 {
		return 5, flanked1, flanked2
	}
	return a + b + 1, flanked1, flanked2
}
