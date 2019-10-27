package play

import (
	"fmt"
	"os"
)

func CanCaptureVertex(coordinate coordinate, goban *[19][19]position, y int8, x int8, player bool) bool {
	minusOne := FindNeighbour(coordinate, y, x, -1)
	one := FindNeighbour(coordinate, y, x, 1)
	two := FindNeighbour(coordinate, y, x, 2)
	if PositionOccupiedByPlayer(one, goban, player) {
		if PositionOccupiedByOpponent(minusOne, goban, player) && PositionUnoccupied(two, goban) {
			return true
		}
		if PositionOccupiedByOpponent(two, goban, player) && PositionUnoccupied(minusOne, goban) {
			return true
		}
	}
	return false
}

func CanCaptureVertices(coordinate coordinate, goban *[19][19]position, player bool) bool {
	var y int8
	var x int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				if CanCaptureVertex(coordinate, goban, y, x, player) == true {
					return true
				}
			}
		}
	}
	return false
}

// returns true if its possible to break the aligned 5
func BreakFive(coordinate coordinate, goban *[19][19]position, y int8, x int8, player bool) bool {
	if CanCaptureVertices(coordinate, goban, player) == true {
		return true
	}
	//move along winning string//////////////
	var multiple int8
	for multiple = 1; multiple < 5; multiple++ {
		neighbour := FindNeighbour(coordinate, y, x, multiple)
		if PositionOccupiedByPlayer(neighbour, goban, player) == true {
			if CanCaptureVertices(neighbour, goban, player) == true {
				return true
			}
		} else {
			break
		}
	}
	for multiple = -1; multiple > -5; multiple-- {
		neighbour := FindNeighbour(coordinate, y, x, multiple)
		if PositionOccupiedByPlayer(neighbour, goban, player) == true {
			if CanCaptureVertices(neighbour, goban, player) == true {
				return true
			}
		} else {
			break
		}
	}
	return false
}

func checkVertexAlignFive(coordinate coordinate, goban *[19][19]position, y int8, x int8, player bool) bool {
	var multiple int8
	var a int8
	var b int8
	for multiple = 1; multiple < 5; multiple++ {
		neighbour := FindNeighbour(coordinate, y, x, multiple)
		if PositionOccupiedByPlayer(neighbour, goban, player) == true {
			a++
		} else {
			break
		}
	}
	for multiple = -1; multiple > -5; multiple-- {
		neighbour := FindNeighbour(coordinate, y, x, multiple)
		if PositionOccupiedByPlayer(neighbour, goban, player) == true {
			b++
		} else {
			break
		}
	}
	if a+b >= 4 {
		return true
	}
	return false
}

func AlignFive(coordinate coordinate, goban *[19][19]position, align5 *align5, player bool) (alignedFive bool) {
	var x int8
	var y int8
	for y = -1; y <= 0; y++ {
		for x = -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				return false
			}
			if checkVertexAlignFive(coordinate, goban, y, x, player) == true {
				if BreakFive(coordinate, goban, y, x, player) == true {
					align5.aligned5 = true
					align5.winner = player
					align5.winmove = coordinate
					return true
				}
				return true
			}
		}
	}
	return false
}

func CaptureTen(g *game) (win bool) {
	if g.capture0 >= 10 {
		return true
	}
	if g.capture1 >= 10 {
		return true
	}
	return false
}

func CheckWin(coordinate coordinate, g *game) { //bool {
	if AlignFive(coordinate, &g.goban, &g.align5, g.player) == true {
		if g.align5.aligned5 == true {
			fmt.Printf("Player %v can win by aligning 5, however the other player can break this alignment by capturing a pair\n", g.player)
		} else {
			fmt.Printf("Player %v wins by aligning 5! final move on position y:%d x:%d\n\n", g.player, coordinate.y, coordinate.x)
			os.Exit(-1) ////// rm for gui, just for test
		}
	}
	if CaptureTen(g) == true {
		fmt.Printf("Player %v wins by capturing 10! final move on position y:%d x:%d\n\n", g.player, coordinate.y, coordinate.x)
		os.Exit(-1) ////// rm for gui, just for test
	}
}
