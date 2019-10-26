package play

import (
	"fmt"
	"math/rand"
	"os"
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
	// move		uint32				// how many moves have been played in total (is this desirable/necessary?)
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

func PlaceRandomIfValid(g *game) (validated bool, coordinate coordinate) { //////// for testing
	coordinate = RandomCoordinate()
	valid, whyInvalid := IsMoveValid(coordinate, g)
	if valid == true {
		PlaceStone(coordinate, g.player, &g.goban)
		return true, coordinate
	} else {
		// return whyInvalid to gui
		fmt.Println(whyInvalid) /////
		return false, coordinate
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

func SwapPlayers(player bool) bool {
	if player == false {
		return true
	} else {
		return false
	}
}

func checkVertexAlignFive(coordinate coordinate, goban *[19][19]position, y int8, x int8, player bool) bool {
	var multiple int8
	var a int8
	var b int8
	a = 0
	b = 0 // necessary?
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
	if a+b >= 5 {
		return true
	}
	return false
}

func AlignFive(coordinate coordinate, g *game) bool {
	var x int8
	var y int8
	for y = -1; y <= 0; y++ {
		for x = -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				return false
			}
			if checkVertexAlignFive(coordinate, &g.goban, y, x, g.player) == true {
				return true
			}
		}
	}
	return false
} //break this alignment by capturing a pair
//or if he has already lost four pairs and the opponent can capture one more, therefore winning by capture.

func CheckWin(coordinate coordinate, g *game) { //bool {
	if AlignFive(coordinate, g) == true {
		fmt.Printf("Win! player: %v. final move on position y:%d x:%d\n\n", g.player, coordinate.y, coordinate.x)
		os.Exit(-1) ////// rm for gui, just for test
		// return true
	}
	// if CaptureTen(g) == true {
	// 	return true
	// }
	// return false
}

func GameLoop(g *game) {
	validated := false
	coordinate := RandomCoordinate() /////
	for i := 0; i < 1000; i++ {      //moves
		validated, coordinate = PlaceRandomIfValid(g)
		fmt.Printf("%v\n", validated)
		if validated == true {
			// Capture
			DumpGoban(&g.goban)
			CheckWin(coordinate, g)
			fmt.Printf("player: %v\n", g.player)
			g.player = SwapPlayers(g.player)
		}
	}
	// check if capture
	//		remove captured
	//		update game.captured struct
	// check win
	//		5 in a row
	// 		all win conditions?
	// update game. struct
	//		player change
	//		moves ++
	//	return err
}

func Play() {
	g := InitializeGame()
	GameLoop(g)

	// /// Test DoubleFree
	// zero := coordinate{0, 0}  /////////
	// PlaceIfValid(zero, g)     /////////
	// three := coordinate{1, 1} /////////
	// PlaceIfValid(three, g)    /////////
	// three = coordinate{3, 5}  /////////
	// PlaceIfValid(three, g)    /////////
	// three = coordinate{3, 4}  /////////
	// PlaceIfValid(three, g)    /////////
	// // g.player = true           //////
	// // three = coordinate{3, 2}  ///////// one of the three-aligned obstructed, therefore next move legal.
	// // PlaceIfValid(three, g)    /////////
	// // g.player = false          /////
	// g.player = true          //////
	// three = coordinate{3, 1} /////////
	// PlaceIfValid(three, g)   /////////
	// g.player = false         /////
	// three = coordinate{3, 3} ///////// Double Three, should be rejected
	// PlaceIfValid(three, g)   ///////// Double Three, should be rejected

	// // launch ebiten. render goban and game stats 		// # do first
	// // GameLoop(g)

	// /// Test Place stone
	// PlaceIfValid(RandomCoordinate(), g)
	// PlaceIfValid(RandomCoordinate(), g)
	// PlaceIfValid(RandomCoordinate(), g)
	// PlaceIfValid(RandomCoordinate(), g)
	// PlaceIfValid(RandomCoordinate(), g)
	// PlaceIfValid(RandomCoordinate(), g)
	// // PlaceStone(RandomCoordinate(), true, &g.goban)  ////////
	// // PlaceStone(RandomCoordinate(), true, &g.goban)  //////
	// // PlaceStone(RandomCoordinate(), true, &g.goban)  ///////
	// // PlaceStone(RandomCoordinate(), false, &g.goban) ////////
	// // PlaceStone(RandomCoordinate(), true, &g.goban)  ////////

	// RemoveStone(zero, &g.goban) /////////

	// coordinate, elapsed := ai(g)
	// PlaceIfValid(coordinate, g) //////
	// fmt.Println(elapsed)        // tme taken by ai //////////

	// gui.RunEbiten()
	// DumpGoban(&g.goban)/////
}
