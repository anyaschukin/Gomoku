package gomoku

// import "fmt"

// func checkNeighbors(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
// 	var a int8
// 	for a = 1; a <= 4; a++ {
// 		neighbour := findNeighbour(coordinate, y, x, a)
// 		if positionOccupiedByOpponent(neighbour, goban, player) == true {
// 			fmt.Printf("Flanked: coordinate = %v, neighbour = %v, player = %v\n", coordinate, neighbour, player)
// 			return true
// 		}
// 	}
// 	return false
// }

// func checkFlanked(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
// 	if checkNeighbors(coordinate, goban, y, x, player) == true || checkNeighbors(coordinate, goban, -y, -x, player) == true {
// 		return true
// 	}
// 	return false
// }

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
	// var flanked bool
	a, flanked1 := measureOpponent(coordinate, goban, y, x, player)
	b, flanked2 := measureOpponent(coordinate, goban, -y, -x, player)
	// if flanked1 == true || flanked2 == true {
	// 	flanked = true
	// }
	if a + b > 4 {
		return 4, flanked1, flanked2
	}
	return a + b, flanked1, flanked2
	// if a+b >= 4 {
	// return 4, flanked
	// } else if a > b {
	// return a, flanked
	// }
	// return b, flanked
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
	// var flanked bool
	a, flanked1 := measurePlayer(coordinate, goban, y, x, player)
	b, flanked2 := measurePlayer(coordinate, goban, -y, -x, player)
	// if flanked1 == true || flanked2 == true {
	// 	flanked = true
	// }
	if a + b + 1 > 5 {
		return 5, flanked1, flanked2
	}
	return a + b + 1, flanked1, flanked2
}
