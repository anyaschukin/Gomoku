package gomoku

import "fmt"

// willCaptureVertex returns true if given coordinate will capture in the next move
func willCaptureDirection(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
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

// willCaptureVertex returns true if given coordinate will be captured in the next move (given opponent who always takes captures)
// func willBeCapturedDirection(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
// 	minusOne := findNeighbour(coordinate, y, x, -1)
// 	two := findNeighbour(coordinate, y, x, 1)
// 	three := findNeighbour(coordinate, y, x, 2)
// 	if positionOccupiedByOpponent(minusOne, goban, player) == true &&
// 		positionOccupiedByPlayer(two, goban, player) == true &&
// 		positionUnoccupied(three, goban) == true {
// 		return true
// 	}
// 	return false
// }

// willCapture returns true if given coordinate will capture (for player) in given direction in the next move
func willCapture(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
	if willCaptureDirection(coordinate, goban, y, x, player) == true || willCaptureDirection(coordinate, goban, -y, -x, player) == true {
		return true
	}
	return false
}

// willBeCaptured returns true if given coordinate will be captured (for opponent) in given direction in the next move
func willBeCaptured(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
	if canBeCapturedVertex(coordinate, goban, y, x, player) == true || canBeCapturedVertex(coordinate, goban, -y, -x, player) == true {
		return true
	}
	return false
}

func captureAttackDefend(coordinate coordinate, goban *[19][19]position, y, x int8, player bool, captures captures) int {
	if willCapture(coordinate, goban, y, x, player) == true {
		if capturedEight(player, captures.capture0, captures.capture1) == true {
			fmt.Printf("Will capture eight! Coordinate: %v, Player: %v\n", coordinate, player)
			return maxInt
		}
		fmt.Printf("Will Capture! Coordinate: %v, Player: %v\n", coordinate, player)
		return 42e13
	} else if willBeCaptured(coordinate, goban, y, x, player) == true {
		if capturedEight(!player, captures.capture0, captures.capture1) == true {
			fmt.Printf("Will be captured 10! Coordinate: %v, Player: %v\n", coordinate, player)
			return -42e15
		}
		fmt.Printf("Will Be Captured! Coordinate: %v, Player: %v\n", coordinate, player)
		return -42e11
	}
	return 0
}

func checkNeighbors(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
	var a int8
	for a = 1; a <= 4; a++ {
		neighbour := findNeighbour(coordinate, y, x, a)
		if positionOccupiedByOpponent(neighbour, goban, player) == true {
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

// checks either side for whose chain it is << REWRITE THIS COMMENT
// func checkNeighbors(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
// 	neighbour1 := findNeighbour(coordinate, y, x, -1)
// 	neighbour2 := findNeighbour(coordinate, y, x, 1)
// 	if positionOccupiedByOpponent(neighbour1, goban, player) == true || positionOccupiedByOpponent(neighbour2, goban, player) == true {
// 		return true
// 	}
// 	return false
// }

// measureChain returns how many stones in a row for given coordinate, axes & player
func measureChain2(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) int8 {
	var length int8
	var multiple int8
	for multiple = 1; multiple <= 5; multiple++ {
		neighbour := findNeighbour(coordinate, y, x, multiple)
		if positionOccupiedByPlayer(neighbour, goban, player) == true {
			length++
		} else {
			break
		}
	}
	return length
}

// chainLength returns the total length of stones aligned running through given a coordinate on a given axe
func chainLength(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) int8 {
	a := measureChain2(coordinate, goban, y, x, player)
	b := measureChain2(coordinate, goban, -y, -x, player)
	return a + b
}

// map[string]bool
// DeepEqual?
