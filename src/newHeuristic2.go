package play

import (
// "fmt"
// "time"
// "os"
)

// the weight of an empty point
// epsilon := 2

// the weights of the adjacent points of influence
// w(k+1) := 2^12, w(k+2):= 2^11, w(k+3) := 2^10, w(k+4) := 2^9

// if a border or a white stone is encountered, the remaining w(k) values in that direction are all set to 1

// the scores of the four directions are combined (by addition) to make up the evaluation score

func coordinateOnBorder(coordinate coordinate) bool {
	if coordinate.y == 0 || coordinate.y == 18 || coordinate.x == 0 || coordinate.x == 18 {
		return true
	}
	return false
}

func weight(z int8) int {
	var influence int

	switch z {
	case 1:
		influence = 1
	case 2:
		influence = 2 ^ 12
	case 3:
		influence = 2 ^ 11
	case 4:
		influence = 2 ^ 10
	case 5:
		influence = 2 ^ 9
	}
	return influence
}

func calcLine(evalAxis int, neighbour coordinate, goban *[19][19]position, player bool, z int8) int {
	epsilon := 2

	if positionOccupied(neighbour, goban) == false { /* if neighbour is empty */
		evalAxis = evalAxis * epsilon
		// fmt.Printf("evalAxis = %v, epsilon = %v\n", evalAxis, epsilon)
	} else if positionOccupiedByPlayer(neighbour, goban, player) == true { /* neighbour is own stone */
		evalAxis = evalAxis * weight(z)
	}
	return evalAxis
}

// calculates the influence of { ownStone, empty spaces, opponentStone/border } at each space in one direction
func lineInfluence(coordinate coordinate, goban *[19][19]position, player bool, y int8, x int8) int {
	var z int8

	evalAxis := 1
	for z = 1; z <= 5; z++ { // for each point
		neighbour := findNeighbour(coordinate, y, x, z)
		if coordinateOnGoban(neighbour) == false {
			continue
		} else if positionOccupiedByOpponent(neighbour, goban, player) == true || coordinateOnBorder(neighbour) == true {
			break
		} else {
			evalAxis = calcLine(evalAxis, neighbour, goban, player, z)
		}
	}
	// fmt.Printf("z = %v\n", z)
	// time.Sleep(300 * time.Millisecond)
	for z = -1; z <= -5; z-- { // for each point
		neighbour := findNeighbour(coordinate, y, x, z)
		if coordinateOnGoban(neighbour) == false {
			continue
		} else if positionOccupiedByOpponent(neighbour, goban, player) == true || coordinateOnBorder(neighbour) == true {
			break
		} else {
			evalAxis = calcLine(evalAxis, neighbour, goban, player, z)
		}
	}
	// fmt.Printf("z = %v, evalAxis = %v\n", z, evalAxis)
	// time.Sleep(300 * time.Millisecond)
	return evalAxis
}

// alignFive returns true if 5 stones are aligned running through given coodinate
func moveEvaluationAlgorithm(coordinate coordinate, goban *[19][19]position, player bool) int {
	var x int8
	var y int8

	eval := 0
	for y = -1; y <= 0; y++ {
		for x = -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				return eval
			}
			eval += lineInfluence(coordinate, goban, player, y, x)
			// evalAxis = lineInfluence(evalAxis, coordinate, goban, player, y, x)
			// eval += evalAxis
		}
	}
	return eval
}
