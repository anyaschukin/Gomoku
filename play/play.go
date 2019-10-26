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
	goban    [19][19]position
	player   bool  // whose move is it? (player 0 - black first)
	capture0 uint8 // capture 10 and win
	capture1 uint8 // capture 10 and win
	// move		uint32				// how many moves have been played in total (is this desirable/necessary?)
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

func CaptureTen(g *game) (win bool, player bool) {
	if g.capture0 >= 10 {
		return true, false
	}
	if g.capture1 >= 10 {
		return true, true
	}
	return false, false
}

func CheckWin(coordinate coordinate, g *game) { //bool {
	if AlignFive(coordinate, g) == true {
		fmt.Printf("Player %v wins by aligning 5! final move on position y:%d x:%d\n\n", g.player, coordinate.y, coordinate.x)
		os.Exit(-1) ////// rm for gui, just for test
		// return true
	}
	captureTen, player := CaptureTen(g)
	if captureTen == true {
		fmt.Printf("Player %v wins by capturing 10! final move on position y:%d x:%d\n\n", player, coordinate.y, coordinate.x)
		os.Exit(-1) ////// rm for gui, just for test
	}
	// return false
}

func CaptureVertex(coordinate coordinate, g *game, y int8, x int8) {
	one := FindNeighbour(coordinate, y, x, 1)
	two := FindNeighbour(coordinate, y, x, 2)
	three := FindNeighbour(coordinate, y, x, 3)
	if PositionOccupiedByOpponent(one, &g.goban, g.player) == true &&
		PositionOccupiedByOpponent(two, &g.goban, g.player) == true &&
		PositionOccupiedByPlayer(three, &g.goban, g.player) == true {
		RemoveStone(one, &g.goban)
		RemoveStone(two, &g.goban)
		fmt.Printf("Capture! player: %v. captured y:%d x:%d & y:%d x:%d\n\n", g.player, one.y, one.x, two.y, two.x) ///
		if g.player == false {
			g.capture0 += 2
		} else {
			g.capture1 += 2
		}
		fmt.Printf("capture0: %d, capture1: %d\n", g.capture0, g.capture1) /////////
		// os.Exit(0) ///////////
	}
}

func Capture(coordinate coordinate, g *game) {
	var x int8
	var y int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				CaptureVertex(coordinate, g, y, x)
			}
		}
	}
}

func GameLoop(g *game) {
	validated := false
	coordinate := RandomCoordinate() /////
	for i := 0; i < 1000; i++ {      //moves
		validated, coordinate = PlaceRandomIfValid(g)
		fmt.Printf("%v\n", validated)
		if validated == true {
			Capture(coordinate, g)
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
