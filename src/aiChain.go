package gomoku

import "fmt"

func checkNeighbors(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
	var a int8
	for a = 1; a <= 4; a++ {
		neighbour := findNeighbour(coordinate, y, x, a)
		if positionOccupiedByOpponent(neighbour, goban, player) == true {
			fmt.Printf("Flanked: coordinate = %v, neighbour = %v, player = %v\n", coordinate, neighbour, player)
			return true
		}
	}
	return false
}

func checkFlanked(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
	if checkNeighbors(coordinate, goban, y, x, player) == true || checkNeighbors(coordinate, goban, -y, -x, player) == true {
		return true
	}
	return false
}

func lengthOpponentChain(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) (int8, bool) {
	var a int8
	var length int8
	for a = 1; a <= 4; a++ {
		neighbour := findNeighbour(coordinate, y, x, a)
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

func lengthDefend(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) (int8, bool) {
	// var a int8
	// var b int8
	var flanked bool
	a, tmp1 := lengthOpponentChain(coordinate, goban, y, x, player)
	b, tmp2 := lengthOpponentChain(coordinate, goban, -y, -x, player)
	if tmp1 == true || tmp2 == true {
		flanked = true
	}
	// fmt.Printf("a = %d, b = %d\n", a, b)
	if a+b >= 4 {
		return 4, flanked
	} else if a > b {
		return a, flanked
	}
	return b, flanked
}

func lengthPlayerChain(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) (int8, bool) {
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

// chainLength returns the total length of stones aligned running through given a coordinate on a given axe
func lengthAttack(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) (int8, bool) {
	// var a int8
	// var b int8
	var flanked bool
	a, tmp1 := lengthPlayerChain(coordinate, goban, y, x, player)
	b, tmp2 := lengthPlayerChain(coordinate, goban, -y, -x, player)
	if tmp1 == true || tmp2 == true {
		flanked = true
	}
	return a + b + 1, flanked
}
