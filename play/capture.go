package play

import (
	"fmt"
	"os"
)

func removeStone(coordinate coordinate, goban *[19][19]position) {
	goban[coordinate.y][coordinate.x].occupied = false
}

func captureVertex(coordinate coordinate, g *game, y int8, x int8) {
	one := FindNeighbour(coordinate, y, x, 1)
	two := FindNeighbour(coordinate, y, x, 2)
	three := FindNeighbour(coordinate, y, x, 3)
	if PositionOccupiedByOpponent(one, &g.goban, g.player) == true &&
		PositionOccupiedByOpponent(two, &g.goban, g.player) == true &&
		PositionOccupiedByPlayer(three, &g.goban, g.player) == true {
		removeStone(one, &g.goban)
		removeStone(two, &g.goban)
		// fmt.Printf("Capture! player: %v. captured y:%d x:%d & y:%d x:%d\n\n", g.player, one.y, one.x, two.y, two.x) ///
		if g.player == false {
			g.capture0 += 2
		} else {
			g.capture1 += 2
		}
		// os.Exit(0) ///////////
	}
}

func Capture(coordinate coordinate, g *game) {
	var x int8
	var y int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				captureVertex(coordinate, g, y, x)
			}
		}
	}
	if g.align5.aligned5 == true { /// check captureTen first: win by align 5 if opponent can not break this alignment by capturing, or if he has already lost four pairs and the opponent can capture one more, therefore winning by capture.
		if PositionOccupiedByPlayer(g.align5.winmove, &g.goban, g.align5.winner) == true &&
			AlignFive(g.align5.winmove, &g.goban, &g.align5, g.align5.winner, g.capture0, g.capture1) == true {
			fmt.Printf("Player %v win by aligning 5.\nThe other player could have broken this alignment by capturing a pair, but they didn't, silly!\nWinning move y:%d x:%d.\n", g.align5.winner, g.align5.winmove.y, g.align5.winmove.x)
			DumpGoban(&g.goban) //////
			os.Exit(-1)         ////// rm, just for test. Return win message to GUI
		}
	}
}
