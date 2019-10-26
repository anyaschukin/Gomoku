package play

import (
	"fmt"
	"math/rand"
	//	lib "Gomoku/golib"
	//	gui "Gomoku/GUI"
)

type coordinate struct {
	y int8
	x int8
}

type position struct { // alternatively uint8 (0 = unocupied), but memory waste
	occupied bool
	player   bool
}

type game struct {
	goban  [19][19]position
	player bool // whose move is it? (player 0 - black first)
	//	capture0	uint8				// capture 10 and win
	//	capture1	uint8				// capture 10 and win
	//	move		uint32				// how many moves have been played in total (is this desirable/necessary?)
	//	pass		bool				// was the last move a pass (if next move pass -> game over)
	//	last0		coordinate			// last move to check ko rule () for player 0 // if player 1 captures multiple stones in next move set to {-1, -1} ko rule need not apply
	//	last1		coordinate			// last move to check ko rule () for player 1

}

// type ai struct { 	/// merge with game struct?
//	aiplayer	bool	// is player 1 human or AI
//	hotseat		bool	// AI player only suggests moves, human must choose move
//	prescience	uint8	// how many moves in advance do we examine
// }

func RandomCoordinate() coordinate {
	x := int8(rand.Intn(19))
	y := int8(rand.Intn(19))
	random := coordinate{y, x}
	return random
}

func PlaceStone(coordinate coordinate, player bool, goban *[19][19]position) {
	goban[coordinate.y][coordinate.x].occupied = true
	goban[coordinate.y][coordinate.x].player = player
}

func RemoveStone(coordinate coordinate, goban *[19][19]position) {
	goban[coordinate.y][coordinate.x].occupied = false
}

func IsMoveValid(coordinate coordinate, g *game) (valid bool, whyInvalid string) {
	if g.goban[coordinate.y][coordinate.x].occupied == true {
		return false, "Position already Occupied"
	}
	if DoubleThree(coordinate, g) == true {
		return false, "Move introduces a forbidden double-three"
	}
	//	other rules?? ko?
	return true, "Valid"
}

func PlaceIfValid(coordinate coordinate, g *game) { /// for human player
	valid, whyInvalid := IsMoveValid(coordinate, g)
	if valid == true {
		PlaceStone(coordinate, g.player, &g.goban)
	} else {
		// return whyInvalid to gui
		fmt.Println(whyInvalid) /////
	}
}

func InitializeGame() *game {
	g := game{}
	// g.player = true ///// rm, just to test
	return &g
}

func DumpGoban(goban *[19][19]position) {
	for y := 0; y < 19; y++ {
		for x := 0; x < 19; x++ {
			if goban[y][x].occupied == true {
				fmt.Printf("\x1B[31m")
			}
			fmt.Printf("{%v\x1B[0m ", goban[y][x].occupied)
			if goban[y][x].occupied == true {
				fmt.Printf(" ")
			}
			color := ""
			if goban[y][x].occupied == true {
				if goban[y][x].player == true {
					color = "\x1B[32m"
				} else {
					color = "\x1B[33m"
				}
			}
			fmt.Printf("%s%v\x1B[0m", color, goban[y][x].player)
			if goban[y][x].player == true {
				fmt.Printf(" ")
			}
			fmt.Printf("} ")
		}
		fmt.Printf("\n")
	}
	// fmt.Println(g.goban) // whole goban //////////
	// fmt.Println(g.goban[0][0])          // one position ///////////
	// fmt.Println(g.goban[0][0].occupied) // one position occupied /////////
	// fmt.Println(g.goban[0][0].player)   // one position player  ///////////
}

func Play() {
	g := InitializeGame()

	zero := coordinate{0, 0}  /////////
	PlaceIfValid(zero, g)     /////////
	three := coordinate{1, 1} /////////
	PlaceIfValid(three, g)    /////////
	three = coordinate{3, 5}  /////////
	PlaceIfValid(three, g)    /////////
	three = coordinate{3, 4}  /////////
	PlaceIfValid(three, g)    /////////
	// g.player = true           //////
	// three = coordinate{3, 2}  ///////// one of the three-aligned obstructed, therefore next move legal.
	// PlaceIfValid(three, g)    /////////
	// g.player = false          /////
	g.player = true          //////
	three = coordinate{3, 1} /////////
	PlaceIfValid(three, g)   /////////
	g.player = false         /////
	three = coordinate{3, 3} ///////// Double Three, should be rejected
	PlaceIfValid(three, g)   ///////// Double Three, should be rejected

	// launch ebiten. render goban and game stats 		// # do first
	// GameLoop(g)

	/// Test Place stone
	PlaceIfValid(RandomCoordinate(), g)
	PlaceIfValid(RandomCoordinate(), g)
	PlaceIfValid(RandomCoordinate(), g)
	PlaceIfValid(RandomCoordinate(), g)
	PlaceIfValid(RandomCoordinate(), g)
	PlaceIfValid(RandomCoordinate(), g)
	// PlaceStone(RandomCoordinate(), true, &g.goban)  ////////
	// PlaceStone(RandomCoordinate(), true, &g.goban)  //////
	// PlaceStone(RandomCoordinate(), true, &g.goban)  ///////
	// PlaceStone(RandomCoordinate(), false, &g.goban) ////////
	// PlaceStone(RandomCoordinate(), true, &g.goban)  ////////

	// RemoveStone(zero, &g.goban) /////////

	// coordinate, elapsed := ai(g)
	// PlaceIfValid(coordinate, g) //////
	// fmt.Println(elapsed)        // tme taken by ai //////////

	// gui.RunEbiten()
	DumpGoban(&g.goban)
}
