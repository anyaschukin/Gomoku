package play

import (
	"fmt"
	"os"
)

// capturedTen returns true if either player has captured ten stones
func capturedTen(g *game) (win bool) {
	if g.capture0 >= 10 || g.capture1 >= 10 {
		return true
	}
	return false
}

// checkVertexAlignFive returns true if 5 stones are aligned running through given coodinate on given axes
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

// AlignFive returns true if 5 stones are aligned running through given coodinate
func AlignFive(coordinate coordinate, goban *[19][19]position, align5 *align5, player bool, capture0 uint8, capture1 uint8) (alignedFive bool) {
	var x int8
	var y int8
	for y = -1; y <= 0; y++ {
		for x = -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				return false
			}
			if checkVertexAlignFive(coordinate, goban, y, x, player) == true {
				if canBreakFive(coordinate, goban, y, x, player) == true {
					align5.break5 = true
				}
				if canWinByCapture(goban, SwapPlayers(player), capture0, capture1) == true {
					align5.capture8 = true
				}
				align5.winner = player
				align5.winmove = coordinate
				return true
			}
		}
	}
	return false
}

// CheckWin checks all win conditions
func CheckWin(coordinate coordinate, g *game) { // return win message (string) to gui
	if capturedTen(g) == true {
		fmt.Printf("Player %v wins by capturing 10! final move on position y:%d x:%d\n\n", g.player, coordinate.y, coordinate.x)
		os.Exit(-1) ////// rm, just for test. Return win message to GUI
	}
	if g.align5.break5 == true {
		if PositionOccupiedByPlayer(g.align5.winmove, &g.goban, g.align5.winner) == true &&
			AlignFive(g.align5.winmove, &g.goban, &g.align5, g.align5.winner, g.capture0, g.capture1) == true {
			fmt.Printf("Player %v win by aligning 5.\nThe other player could have broken this alignment by capturing a pair, but they didn't, silly!\nWinning move y:%d x:%d.\n", g.align5.winner, g.align5.winmove.y, g.align5.winmove.x)
			os.Exit(-1) ////// rm, just for test. Return win message to GUI
		}
		g.align5.break5 = false
	}
	if g.align5.capture8 == true {
		fmt.Printf("Player %v win by aligning 5.\nThe other player could have won by capturing ten, but they didn't, silly!\nWinning move y:%d x:%d.\n", g.align5.winner, g.align5.winmove.y, g.align5.winmove.x)
		os.Exit(-1) ////// rm, just for test. Return win message to GUI
	}
	if AlignFive(coordinate, &g.goban, &g.align5, g.player, g.capture0, g.capture1) == true {
		if g.align5.break5 == true {
			fmt.Printf("Player %v has aligned 5, however the other player can break this alignment by capture\n", g.player)
		} else if g.align5.capture8 == true {
			fmt.Printf("Player %v has aligned 5, however the other player can win by capturing a pair\n", g.player)
		} else {
			fmt.Printf("Player %v wins by aligning 5! final move on position y:%d x:%d\n\n", g.player, coordinate.y, coordinate.x)
			os.Exit(-1) ////// rm, just for test. Return win message to GUI
		}
	}
}
