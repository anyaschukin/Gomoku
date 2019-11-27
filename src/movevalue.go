package play

// func defend against threat
// func play offensive

// struct offensiveplayer bool (player plays either offensively or defensively)
// offensive = 2 offensive moves in a row. Turns off after 1 defensive move?

// fully-open double 3s are illegal

// store coordinates of 2,3,4 patterns in struct

// check align 2,3,4
// check spaced align 2,3,4 (or just 4?)
// check flanked

// calculate score for one move or for whole board?

// checkVertexAlign returns true if 2 or more stones are aligned running through given coodinate on given axes
func checkVertexAlign(coordinate coordinate, goban *[19][19]position, y int8, x int8, player bool, flanked int8, space int8) (int, int8, int8) {
	var multiple int8
	var a int
	var b int
	for multiple = 1; multiple < 5; multiple++ {
		neighbour := findNeighbour(coordinate, y, x, multiple)
		if positionOccupiedByPlayer(neighbour, goban, player) == true {
			a++
		} else if positionOccupiedByOpponent(neighbour, goban, player) == true {
			flanked++
			break
		} else if space == 0 {
			space++
		} else {
			break
		}
	}
	for multiple = -1; multiple > -5; multiple-- {
		neighbour := findNeighbour(coordinate, y, x, multiple)
		if positionOccupiedByPlayer(neighbour, goban, player) == true {
			b++
		} else if positionOccupiedByOpponent(neighbour, goban, player) == true {
			flanked++
			break
		} else if space == 0 {
			space++
		} else {
			break
		}
	}
	// save coordinates of 2,3,4 ?
	if a+b >= 2 && flanked <= 1 {
		return a+b, flanked, space
	}
	return 0, flanked, space
}

// offensive move vs. defensive move... Total the score and select whichever is best?
// maximize your align
// break your opponent's align
// capture opponent's 2s
// defend against capture

// NOTE: Is there a way for use to use this to GENERATE boards only if they align 2,3,4??
// as in... is it necessary to generate a board for every position? Or can we simulate them, then generate and store in struct with value?

// CheckAlign returns true if 2 or more stones are aligned running through given coodinate and assigns a value to that board
func checkAlign(coordinate coordinate, goban *[19][19]position, player bool) int {
	var x		int8
	var y		int8
	var align	int
	var flanked int8
	var space	int8
	var value	int
	for y = -1; y <= 0; y++ {
		for x = -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				return value
			}
			align, flanked, space = checkVertexAlign(coordinate, goban, y, x, player, flanked, space)
			switch align {
				case 5:
					if space != 0 {
						value = align*10
					}
				case 4: 
					value = align*25	// open 4
					if flanked != 0 || space != 0 {
						value -= 30
					}
				case 3:
					value = align*20	// open 3
					// if flanked != 0 || space != 0 {
						// value -= 20
					// }
				case 2:
					value = align*15	// open 2
					// if flanked != 0 {
						// value = -50
					// }
			}
			return value
		}
	}
	return value
}

func valueBoard(goban *[19][19]position, player bool) int {
	var y		int8
	var x		int8
	var value	int
	for y = 0; y < 19; y++ {
		for x = 0; x < 19; x++ {
			coordinate := coordinate{y, x}
			value += checkAlign(coordinate, goban, player)
			// node.Value += value
		}
	}
	return value
}