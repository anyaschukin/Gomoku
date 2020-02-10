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

func threatCaptureDefend(coordinate coordinate, goban *[19][19]position, y, x int8, player bool, length int8, captures *captures) int {
	// dumpGobanBlank(goban)
	fmt.Printf("Hello\n")

	switch length {
	case 5:
		return (maxInt - 1000)
	case 4:
		return 42e12
	case 3:
		return 42e10
	case 2:
		return 42e7
	}
	return -1
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
	// var tmp int

	evalAxis := 1
	for a = 1; a <= 4; a++ {
		neighbour := findNeighbour(coordinate, y, x, a)
		if coordinateOnGoban(neighbour) == false { //not sure this is necessary
			break
		}
		if positionOccupiedByOpponent(neighbour, goban, player) == true || coordinateOnBorder(neighbour) == true {
			evalAxis += int(a)
			break
		} else {
			evalAxis += calcLine(evalAxis, neighbour, goban, player, a)
		}
	}
	for b = -1; b >= -4; b-- {
		neighbour := findNeighbour(coordinate, y, x, b)
		if coordinateOnGoban(neighbour) == false {
			break
		}
		if positionOccupiedByOpponent(neighbour, goban, player) == true || coordinateOnBorder(neighbour) == true {
			evalAxis += int(b)
			break
		} else {
			evalAxis += calcLine(evalAxis, neighbour, goban, player, b)
		}
	}
	return evalAxis
}

// evaluateMove checks for alignments/captures along each vertex for one move, and returns a score for that move
func evaluateMove(coordinate coordinate, goban *[19][19]position, player bool, captures captures) int {
	var x int8
	var y int8

	eval := 0
	for y = -1; y <= 0; y++ {
		for x = -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				return eval
			}
			if willCaptureDirection(coordinate, goban, y, x, player) == true {
				// fmt.Printf("Can Capture\n")
				if canWinByCapture(goban, player, captures.capture0, captures.capture1) == true {
					return maxInt
				}
				eval += 42e11
			}
			length := checkTotalChainLength(coordinate, goban, y, x, player)
			tmp := threatCaptureDefend(coordinate, goban, y, x, player, length, &captures)
			if tmp != -1 {
				eval += tmp
			}
			eval += lineInfluence(coordinate, goban, player, y, x, &captures)
		}
	}
	return eval
}
