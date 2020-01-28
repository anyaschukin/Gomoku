package gomoku

import (
	"fmt"
	"math"
	// "time"
	// "os"
)

// if a border or a white stone is encountered, the remaining w(k) values in that direction are all set to 1
// the scores of the four directions are combined (by addition) to make up the evaluation score

/* the weight of an empty point */
const epsilon = 2

/* if this move captures a 2-in-a-row */
// const captureTwo = 42e8

/* defend against or break a 3-in-a-row */
// const defendThree = 42e11

/* defend against or break a 4-in-a-row */
// const defendFour = 42e12

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

//remove this function
// func measureChain2(coordinate coordinate, goban *[19][19]position, y, x int8, player bool, myChain bool) (bool, int8) {
// 	var i int8

// 	// if neighbor_one is me, MyChain
// 	if positionOccupiedByPlayer(coordinate, goban, player) == true {
// 		myChain = true
// 		for i = 1; i <= 4; i++ {
// 			neighbour := findNeighbour(coordinate, y, x, i)
// 			if positionOccupiedByPlayer(neighbour, goban, player) == false {
// 				break
// 			}
// 		}
// 	} else if positionOccupiedByOpponent(coordinate, goban, player) == true {
// 		myChain = false
// 		for i = 1; i <= 4; i++ {
// 			neighbour := findNeighbour(coordinate, y, x, i)
// 			if positionOccupiedByOpponent(neighbour, goban, player) == false {
// 				break
// 			}
// 		}
// 	}
// 	return myChain, i
// }

func threatCaptureDefend(neighbour coordinate, goban *[19][19]position, y, x int8, player bool, captures *captures) int {
	var myChain bool // replace this variable name
	var length int8

	// myChain, length = measureChain2(coordinate, goban, y, x, player, myChain)
	if positionOccupied(neighbour, goban) == true {
		if positionOccupiedByPlayer(neighbour, &g.goban, player) == true {
			myChain = true
			length = measureChain(neighbour, goban, y, x, player)
		} else {
			length = measureChain(neighbour, goban, y, x, !player)
		}
	}

	switch {
	case length == 4:
		fmt.Printf("Four\n")
		return 42e12
	case length == 3:
		fmt.Printf("Three\n")
		return 42e11
	case length == 2:
		if myChain == false && canWinByCapture(goban, player, captures.capture0, captures.capture1) == true {
			fmt.Printf("Two Win by Capture\n")
			return maxInt
		} else if myChain == false && canCapture(neighbour, goban, player) == true {
			// if capture, then need to modify captures to reflect this
			fmt.Printf("Two Capture: length = %d, myChain = %v\n", length, myChain)
			return 42e10 // how should I weight this differently?
		} else {
			fmt.Printf("Two\n")
			return 42e8
		}

	// myChain = true
	// switch {
	// case length == 4:
	// 	fmt.Printf("Four\n")
	// 	return 42e12
	// case length == 3:
	// 	fmt.Printf("Three\n")
	// 	return 42e11
	// case length == 2, myChain == false, canWinByCapture(goban, player, captures.capture0, captures.capture1) == true:
	// 	fmt.Printf("Two Win by Capture\n")
	// 	return maxInt
	// case length == 2, myChain == false, canCapture(neighbour, goban, player) == true:
	// 	// if capture, then need to modify captures to reflect this
	// 	fmt.Printf("Two Capture: length = %d, myChain = %v\n", length, myChain)
	// 	return 42e9 // how should I weight this differently?
	// case length == 2:
	// 	fmt.Printf("Two\n")
	// 	return 42e8
	}
	return -1
}

// func defend(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) int8 {
// 	one := findNeighbour(coordinate, y, x, 1)
// 	two := findNeighbour(coordinate, y, x, 2)
// 	three := findNeighbour(coordinate, y, x,  3)
// 	four := findNeighbour(coordinate, y, x, 4)
// 	if positionOccupiedByOpponent(one, goban, player) == true &&
// 		positionOccupiedByOpponent(two, goban, player) == true &&
// 		positionOccupiedByOpponent(three, goban, player) == true {
// 		if positionOccupiedByOpponent(four, goban, player) == true {
// 			return 4
// 		}
// 		return 3
// 	}
// 	return 0
// }

// func canCapture2(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
// 	one := findNeighbour(coordinate, y, x, 1)
// 	two := findNeighbour(coordinate, y, x, 2)
// 	three := findNeighbour(coordinate, y, x, 3)
// 	if positionOccupiedByOpponent(one, goban, player) == true &&
// 		positionOccupiedByOpponent(two, goban, player) == true &&
// 		positionOccupiedByPlayer(three, goban, player) == true {
// 		return true
// 	}
// 	return false
// }

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
func lineInfluence(coordinate coordinate, goban *[19][19]position, player bool, y int8, x int8, captures *captures) int {
	var a int8
	var b int8
	var tmp int

	evalAxis := 1
	for a = 1; a <= 4; a++ {
		neighbour := findNeighbour(coordinate, y, x, a)
		if coordinateOnGoban(neighbour) == false {
			break
		}
		tmp = threatCaptureDefend(neighbour, goban, y, x, player, captures)
		if tmp != -1 {
			evalAxis = tmp
			return evalAxis
			// break
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
		tmp = threatCaptureDefend(neighbour, goban, y, x, player, captures)
		if tmp != -1 {
			evalAxis = tmp
			return evalAxis
			// break
		} else if positionOccupiedByOpponent(neighbour, goban, player) == true || coordinateOnBorder(neighbour) == true {
			evalAxis += int(b)
			break
		} else {
			evalAxis += calcLine(evalAxis, neighbour, goban, player, b)
		}

		// d := defend(coordinate, goban, y, x, player)
		// if d == 3 { // attackThree
		// 	return defendThree
		// } else if d == 4 {
		// 	return defendFour
		// } else if canCapture2(coordinate, goban, -y, -x, player) == true {
		// 	evalAxis *= captureTwo
		// 	break
		// }
	}
	// fmt.Printf("evalAxis = %d\n", evalAxis)
	return evalAxis
}

// alignFive returns true if 5 stones are aligned running through given coordinate
func evaluateMove(coordinate coordinate, goban *[19][19]position, player bool, captures captures) int {
	var x int8
	var y int8

	eval := 0
	for y = -1; y <= 0; y++ {
		for x = -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				return eval
			}
			eval += lineInfluence(coordinate, goban, !player, y, x, &captures)
		}
	}
	return eval
}
