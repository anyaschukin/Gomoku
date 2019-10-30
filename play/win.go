package play

import (
	"fmt"
	"os"
)

func captureTen(g *game) (win bool) {
	if g.capture0 >= 10 || g.capture1 >= 10 {
		return true
	}
	return false
}

func canBeCapturedVertex(coordinate coordinate, goban *[19][19]position, y int8, x int8, player bool) bool {
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

func canBeCapturedVertices(coordinate coordinate, goban *[19][19]position, player bool) bool {
	var y int8
	var x int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				if canBeCapturedVertex(coordinate, goban, y, x, player) == true {
					return true
				}
			}
		}
	}
	return false
}

// returns true if its possible to break the aligned 5
func canBreakFive(coordinate coordinate, goban *[19][19]position, y int8, x int8, player bool) bool {
	if canBeCapturedVertices(coordinate, goban, player) == true {
		return true
	}
	//move along winning string//////////////
	var multiple int8
	var a int8
	var b int8
	for multiple = 1; multiple < 5; multiple++ {
		neighbour := FindNeighbour(coordinate, y, x, multiple)
		if PositionOccupiedByPlayer(neighbour, goban, player) == true &&
			canBeCapturedVertices(neighbour, goban, player) == false {
			a++
		} else {
			break
		}
	}
	for multiple = -1; multiple > -5; multiple-- {
		neighbour := FindNeighbour(coordinate, y, x, multiple)
		if PositionOccupiedByPlayer(neighbour, goban, player) == true &&
			canBeCapturedVertices(neighbour, goban, player) == false {
			b++
		} else {
			break
		}
	}
	if a+b >= 4 {
		return false
	}
	return true
}

// capturedEight returns true if given player has already captured 8
func capturedEight(player bool, capture0 uint8, capture1 uint8) bool {
	if player == false {
		if capture0 >= 8 {
			return true
		}
	} else {
		if capture1 >= 8 {
			return true
		}
	}
	return false
}

func canCaptureVertex(coordinate coordinate, goban *[19][19]position, y int8, x int8, player bool) bool {
	one := FindNeighbour(coordinate, y, x, 1)
	two := FindNeighbour(coordinate, y, x, 2)
	three := FindNeighbour(coordinate, y, x, 3)
	if PositionOccupiedByOpponent(one, goban, player) == true &&
		PositionOccupiedByOpponent(two, goban, player) == true &&
		PositionUnoccupied(three, goban) == true {
		// fmt.Printf("Capture possible! player: %v. can capture y:%d x:%d & y:%d x:%d\n\n", g.player, one.y, one.x, two.y, two.x) ///
		return true
	}
	return false
}

func canCapture(coordinate coordinate, goban *[19][19]position, player bool) bool {
	var y int8
	var x int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				if canCaptureVertex(coordinate, goban, y, x, player) == true {
					return true
				}
			}
		}
	}
	return false
}

// captureAvailable returns true if given player can capture in the next move (iterate entire goban, check if capture possible for each positon)
func captureAvailable(goban *[19][19]position, player bool) bool {
	var y int8
	var x int8
	for y = 0; y < 19; y++ {
		for x = 0; x < 19; x++ {
			coordinate := coordinate{y, x}
			if PositionOccupiedByPlayer(coordinate, goban, player) == true {
				if canCapture(coordinate, goban, player) == true {
					return true
				}
			}
		}
	}
	return false
}

// canWinByCapture returns true if is it possible for the opponent to win by capturing 10. (have they already captured 8, and is there an available capture move)
func canWinByCapture(goban *[19][19]position, player bool, capture0 uint8, capture1 uint8) bool {
	if capturedEight(player, capture0, capture1) == true &&
		captureAvailable(goban, player) == true {
		return true
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

func CheckWin(coordinate coordinate, g *game) { //bool {
	if captureTen(g) == true {
		fmt.Printf("Player %v wins by capturing 10! final move on position y:%d x:%d\n\n", g.player, coordinate.y, coordinate.x)
		os.Exit(-1) ////// rm, just for test. Return win message to GUI
	}
	if g.align5.break5 == true { /// check captureTen first: win by align 5 if opponent can not break this alignment by capturing, or if he has already lost four pairs and the opponent can capture one more, therefore winning by capture.
		if PositionOccupiedByPlayer(g.align5.winmove, &g.goban, g.align5.winner) == true &&
			AlignFive(g.align5.winmove, &g.goban, &g.align5, g.align5.winner, g.capture0, g.capture1) == true {
			fmt.Printf("Player %v win by aligning 5.\nThe other player could have broken this alignment by capturing a pair, but they didn't, silly!\nWinning move y:%d x:%d.\n", g.align5.winner, g.align5.winmove.y, g.align5.winmove.x)
			os.Exit(-1) ////// rm, just for test. Return win message to GUI
		}
	}
	if g.align5.capture8 == true {
		if PositionOccupiedByPlayer(g.align5.winmove, &g.goban, g.align5.winner) == true &&
			AlignFive(g.align5.winmove, &g.goban, &g.align5, g.align5.winner, g.capture0, g.capture1) == true {
			fmt.Printf("Player %v win by aligning 5.\nThe other player could have won by capturing ten, but they didn't, silly!\nWinning move y:%d x:%d.\n", g.align5.winner, g.align5.winmove.y, g.align5.winmove.x)
			os.Exit(-1) ////// rm, just for test. Return win message to GUI
		}
	}
	if AlignFive(coordinate, &g.goban, &g.align5, g.player, g.capture0, g.capture1) == true {
		if g.align5.break5 == true {
			fmt.Printf("Player %v can win by aligning 5, however the other player can break this alignment by capturing a pair\n", g.player)
		} else if g.align5.capture8 == true {
			fmt.Printf("Player %v can win by aligning 5, however the other player can win by capturing a pair\n", g.player)
		} else {
			fmt.Printf("Player %v wins by aligning 5! final move on position y:%d x:%d\n\n", g.player, coordinate.y, coordinate.x)
			os.Exit(-1) ////// rm, just for test. Return win message to GUI
		}
	}
}
