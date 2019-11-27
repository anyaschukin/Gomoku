package play

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

// calculates the influence of { ownStone, empty spaces, opponentStone/border } at each space in one direction
func lineInfluence(evalAxis int, coordinate coordinate, goban *[19][19]position, player bool, y int8, x int8) int {
	var z int8

	epsilon := 2
	for z = 1; z <= 5; z++ { // for each point
		neighbour := findNeighbour(coordinate, y, x, z)
		if positionOccupiedByOpponent(neighbour, goban, player) == true || coordinateOnBorder(neighbour) == true {
			break
		} else if positionOccupied(neighbour, goban) == false { /* if neighbour is empty */
			evalAxis = evalAxis * epsilon
		} else if positionOccupiedByPlayer(neighbour, goban, player) == true { /* neighbour is own stone */
			evalAxis = evalAxis * weight(z)
		}
	}
	return evalAxis
}

func moveEvaluationAlgorithm(coordinate coordinate, goban *[19][19]position, player bool) int {
	var y int8
	var x int8

	eval := 0
	evalAxis := 1
	for y = 1; y <= 4; y++ { // for four directions (vertical, horizontal, diagonal1, diagonal2)
		for x = 1; x <= 5; x++ { // for half of a line, in one direction
			evalAxis = lineInfluence(evalAxis, coordinate, goban, player, y, x)
			eval += evalAxis
		}
		for x = -1; x <= -5; x-- { // for half of a line, in one direction
			evalAxis = lineInfluence(evalAxis, coordinate, goban, player, y, x)
			eval += evalAxis
		}
	}
	return eval
}

// for each half of a line (a goes 5 in one direction, b goes 5 in the other direction)
// for k = 1; k <= 5; k++ { // for each point
// 	neighbour := findNeighbour(coordinate, y, x, k)
// 	if positionOccupiedByOpponent(neighbour, goban, player) == true /*|| border*/ {
// 		break
// 	} else if positionOccupied(coordinate, goban) == false { /* if coordinate is empty */
// 		evalAxis = evalAxis * epsilon
// 	} else if positionOccupiedByPlayer(coordinate, goban) == true { /* coordinate is own stone */
// 		evalAxis = evalAxis * weight(k)
// 	}
// }
