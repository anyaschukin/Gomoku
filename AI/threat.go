package ai

import (
	"fmt"
	play "Gomoku.play"
)

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
func checkVertexAlign(coordinate coordinate, Goban *[19][19]position, y int8, x int8, Player bool, flanked int8, space int8) int8, int8, int8 {
	var multiple int8
	var a int8
	var b int8
	for multiple = 1; multiple < 5; multiple++ {
		neighbour := FindNeighbour(coordinate, y, x, multiple)
		if PositionOccupiedByPlayer(neighbour, Goban, Player) == true {
			a++
		} else if PositionOccupiedByOpponent(neighbour, Goban, Player) == true {
			flanked++
			break
		} else if space == 0 {
			space++
		} else {
			break
		}
	}
	for multiple = -1; multiple > -5; multiple-- {
		neighbour := FindNeighbour(coordinate, y, x, multiple)
		if PositionOccupiedByPlayer(neighbour, Goban, Player) == true {
			b++
		} else if PositionOccupiedByOpponent(neighbour, Goban, Player) == true {
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

// NOTE: Is there a way for use to use this to GENERATE boards only if they align 2,3,4??
// as in... is it necessary to generate a board for every position? Or can we simulate them, then generate and store in struct with value?

// CheckAlign returns true if 2 or more stones are aligned running through given coodinate and assigns a value to that board
func checkAlign(coordinate coordinate, Goban *[19][19]position, Player bool, value int8) int8 {
	var x		int8
	var y		int8
	var align	int8
	var flanked int8
	var space	int8
	for y = -1; y <= 0; y++ {
		for x = -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				return value
			}
			align, flanked, space := checkVertexAlign(coordinate, Goban, y, x, Player, flanked, space)
			switch align {
				case 5:
					if space {
						value := align*10
					}
				case 4: 
					value := align*25	// open 4
					if flanked || space {
						value -= 30
					}
				case 3:
					value := align*20	// open 3
					if flanked || space {
						value -= 20
					}
				case 2:
					value := align*15	// open 2
					if flanked {
						value = -50
					}
			}
			return value
			}
		}
	}
	return value
}

func valueBoard(Goban *[19][19]position, Player bool) {
	var y		int8
	var x		int8
	for y = 0; y < 19; y++ {
		for x = 0; x < 19; x++ {
			coordinate := coordinate{y, x}
			value := checkAlign(coordinate, goban, player, value)
			node.Value += value
			}
		}
	}
	return false
}