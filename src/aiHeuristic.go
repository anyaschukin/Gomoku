package gomoku

import (
	"fmt"
	"math"
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

// chainAttackDefend returns a score for aligning 5, 4, 3, or 2 stones
func chainAttackDefend(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) int {
	// dumpGobanBlank(goban)

	// if opponent will capture

	// if willCapture(coordinate, goban, y, x, player) == true {
	// 	fmt.Printf("Will capture - player = %v\n", player)
	// 	// check that canWinByCapture works
	// 	if canWinByCapture(goban, player, captures.capture0, captures.capture1) == true {
	// 		return maxInt
	// 	}
	// 	return 42e11
	// }

	// capt := captureOrBeCaptured(coordinate, goban, y, x, player)
	// if capt != 0 {
	// 	fmt.Printf("capt = %d\n", capt)
	// 	return capt
	// }

	// heuristic prioritizes blocking opponent's 4 over aligning own 5
	defend := checkNeighbors(coordinate, goban, y, x, player)
	switch defend {
	case true:
		length := chainLength(coordinate, goban, y, x, !player)
		if length >= 2 {
			// fmt.Printf("opponent's length = %d\n", length)
		}
		switch length {
		case 4:
			// fmt.Printf("opponent 4, player = %v\n", player)
			return 42e14
		case 3:
			// fmt.Printf("opponent 3, player = %v\n", player)
			return 42e10
		}
	case false:
		length := chainLength(coordinate, goban, y, x, !player)
		length++
		if length >= 2 {
			// fmt.Printf("player's length = %d\n", length)
		}
		switch length {
		case 5:
			return 42e14
			// return (maxInt - 1000)
		case 4:
			// PLAYS 4 even if Flanked on both sides
			// fmt.Printf("player 4, player = %v\n", player)
			return 42e12
		case 3:
			// fmt.Printf("player 3, player = %v\n", player)
			return 42e10
		case 2:
			// fmt.Printf("player 2, player = %v\n", player)
			return 42e7
		}
	}
	return 0
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
			// captureAttackDefend
			// chainAttackDefend
			capt := captureAttackDefend(coordinate, goban, y, x, player, captures)
			if capt == maxInt {
				return capt
			}
			eval += capt
			// if capt != 0 {
			// fmt.Printf("capt = %d\n", capt)
			// return capt
			// }
			// make this either-or
			tmp := chainAttackDefend(coordinate, goban, y, x, player)
			if tmp == 0 {
				tmp = lineInfluence(coordinate, goban, player, y, x, &captures)
			}
			eval += tmp
		}
	}
	return eval
}

// TO DO
//  replace aiChain's canBeCapturedVertex (line 53) with willBeCapturedVertex
//  if willCapture && willBeCapturedVertex == true?? How to score?
// - heuristic does not prioritize winning 10th capture, instead prioritzes 3-align

//  adding capt + tmp + eval == rollover :( ... and then stupid moves!
// 9223372036854775807
// 4.204200000000000000
