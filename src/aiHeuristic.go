package gomoku

import (
	"math"
	// "time"
	// "os"
)

// if a border or a white stone is encountered, the remaining w(k) values in that direction are all set to 1
// the scores of the four directions are combined (by addition) to make up the evaluation score

/* the weight of an empty point */
const epsilon = 2

/* if this move captures a 2-in-a-row */
const captureTwo = 42e8

/* defend against or break a 3-in-a-row */
const defendThree = 42e11

/* defend against or break a 4-in-a-row */
const defendFour = 42e12

// the weights of the adjacent points of influence
// w(k+1) := 2^12, w(k+2):= 2^11, w(k+3) := 2^10, w(k+4) := 2^9
func weight(z int8) int {
	var influence float64

	switch z {
	case 1:
		influence = math.Pow(2, 12)
	case 2:
		influence = math.Pow(2, 11)
	case 3:
		influence = math.Pow(2, 10)
	case 4:
		influence = math.Pow(2, 9)
	}
	return int(influence)
}

// threat, capture, defend

func measureChain(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) (bool, int8) {
	var myChain bool  // replace this variable name
	var i int8

	// if neighbor_one is me, MyChain
	if positionOccupiedByPlayer(coordinate, goban, player) == true {
		myChain = true
		for i = 1; i <= 4; i++ {
			neighbour := findNeighbour(coordinate, y, x, i)
			if positionOccupiedByPlayer(neighbour, goban, player) == false {
				break
			}
		}
	} else if positionOccupiedByOpponent(coordinate, goban, player) == true {
		myChain = false
		for i = 1; i <= 4; i++ {
			neighbour := findNeighbour(coordinate, y, x, i)
			if positionOccupiedByOpponent(neighbour, goban, player) == false {
				break
			}
		}
	}
	return myChain, i
}

func threatCaptureDefend(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) {
	if positionOccupied(coordinate, goban) == true {
		myChain, length := measureChain(coordinate, goban, y, x, player)
	}

	// g.capture0, g.capture1
	switch length {
	// need to include a myChain condition here
	case 2 && canWinByCapture() == true: // does this work or do I ALSO need if canCapture() in here?
		return Win
	// need to include a myChain condition here
	case 2 && canCapture() == true:
		return Two // how should I weight this differently?
	case 2:
		return Two
	case 3:
		return Three
	case 4:
		return Four
	}

}

func defend(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) int8 {
	one := findNeighbour(coordinate, y, x, 1)
	two := findNeighbour(coordinate, y, x, 2)
	three := findNeighbour(coordinate, y, x, 3)
	four := findNeighbour(coordinate, y, x, 4)
	if positionOccupiedByOpponent(one, goban, player) == true &&
		positionOccupiedByOpponent(two, goban, player) == true &&
		positionOccupiedByOpponent(three, goban, player) == true {
		if positionOccupiedByOpponent(four, goban, player) == true {
			return 4
		}
		return 3
	}
	return 0
}

func canCapture2(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
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

func coordinateOnBorder(coordinate coordinate) bool {
	if coordinate.y == 0 || coordinate.y == 18 || coordinate.x == 0 || coordinate.x == 18 {
		return true
	}
	return false
}

func calcLine(evalAxis int, neighbour coordinate, goban *[19][19]position, player bool, z int8) int {
	if positionOccupied(neighbour, goban) == false { /* if neighbour is empty */
		evalAxis *= epsilon
	} else if positionOccupiedByPlayer(neighbour, goban, player) == true { /* neighbour is own stone */
		evalAxis *= weight(z)
	}
	return evalAxis
}

// calculates the influence of { ownStone, empty spaces, opponentStone, border } at each space in one direction
func lineInfluence(coordinate coordinate, goban *[19][19]position, player bool, y int8, x int8) int {
	var a int8
	var b int8

	evalAxis := 1
	for a = 1; a <= 4; a++ {
		neighbour := findNeighbour(coordinate, y, x, a)
		if coordinateOnGoban(neighbour) == false {
			break
		}
		tmp := threatCaptureDefend(neighbour, goban, y, x, player) 
		if tmp != 0 {
			evalAxis = tmp
			break
		} else if positionOccupiedByOpponent(neighbour, goban, player) == true || coordinateOnBorder(neighbour) == true {
			evalAxis += int(a)
			break
		} else {
			evalAxis += calcLine(evalAxis, neighbour, goban, player, a)
		}

		// d := defend(coordinate, goban, y, x, player)
		// if d == 3 { // attackThree
		// 	return defendThree
		// } else if d == 4 {
		// 	return defendFour
		// } else if canCapture2(coordinate, goban, y, x, player) == true { //
		// 	evalAxis *= captureTwo
		// 	break
		// } 
	}
	// if evalAxis == maxInt ... RETURN!
	for b = -1; b >= -4; b-- {
		neighbour := findNeighbour(coordinate, y, x, b)
		if coordinateOnGoban(neighbour) == false {
			break
		}
		d := defend(coordinate, goban, y, x, player)
		if d == 3 { // attackThree
			return defendThree
		} else if d == 4 {
			return defendFour
		} else if canCapture2(coordinate, goban, -y, -x, player) == true {
			evalAxis *= captureTwo
			break
		} else if positionOccupiedByOpponent(neighbour, goban, player) == true || coordinateOnBorder(neighbour) == true {
			evalAxis += int(b)
			break
		} else {
			evalAxis += calcLine(evalAxis, neighbour, goban, player, b)
		}
	}
	return evalAxis
}

// alignFive returns true if 5 stones are aligned running through given coordinate
func evaluateMove(coordinate coordinate, goban *[19][19]position, player bool) int {
	var x int8
	var y int8

	eval := 0
	for y = -1; y <= 0; y++ {
		for x = -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				return eval
			}
			eval += lineInfluence(coordinate, goban, !player, y, x)
			// eval -= lineInfluence(coordinate, goban, !player, y, x)
		}
	}
	return eval
}
