package play

// the weight of an empty poin
// epsilon := 2 

// the weights of the adjacent points of influence
// w(k+1) := 2^12, w(k+2):= 2^11, w(k+3) := 2^10, w(k+4) := 2^9

// if a border or a white stone is encountered, the remaining w(k) values in that direction are all set to 1


// the scores of the four directions are combined (by addition) to make up the evaluation score

func moveEvaluationAlgorithm(coordinate coordinate, goban *[19][19]position, player bool) int {
	eval := 0
	evalAxis := 1

	for j = 1; j <= 4; j++ {			// for four directions (vertical, horizontal, diagonal1, diagonal2)
		for /* l...a,b */ {						// for each half of a line (a goes 5 in one direction, b goes 5 in the other direction)
			for k = 1; k <= 5; k++ {	// for each point
				if positionOccupiedByOpponent(neighbour, goban, player) == true || border {
					break
				} else if /*coordinate is empty */ {
					evalAxis = evalAxis * epsilon
				} else if /*coordinate is own stone*/ {
					evalAxis = evalAxis * weight(k)
				}
			}
			eval += evalAxis
		}
	}
	return eval
}

