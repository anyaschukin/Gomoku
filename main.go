package main

import (
	"fmt"
	"math/rand"
	"time"
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

func InitializeGame() *game {
	g := game{}
	// g.player = true ///// rm, just to test
	return &g
}

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

// func GameLoop(g) {
// if AI:
//		suggest move
// if human or hotseat:
//		listen for mouse click 						//###### do first
//			find position/pass/new/exit clicked 		//###### do first
//			if pass, double pass end?
//			if reset, reset game with new config
// 		check if position is valid (if human, assume ai has aleady checked)
//			occupied?
//			rules
//				ko
//				double-three
// PlaceStone(coordinate, true, &g.goban)
// check if capture
//		remove captured
//		update game.captured struct
// check win
//		5 in a row
// 		all win conditions?
// update game. struct
//		player change
//		moves ++
// re-render ebiten with updated goban and stats
//	return err
// }

func CoordinateOnGoban(coordinate coordinate) (onGoban bool) {
	if coordinate.y < 0 || coordinate.y > 18 || coordinate.x < 0 || coordinate.x > 18 {
		return false
	}
	return true
}

func PositionOccupied(coordinate coordinate, goban *[19][19]position) (occupied bool) {
	if goban[coordinate.y][coordinate.x].occupied == true {
		return true
	}
	return false
}

func SamePlayer(coordinate coordinate, goban *[19][19]position, player bool) (samePlayer bool) {
	if goban[coordinate.y][coordinate.x].player == player {
		return true
	}
	return false
}

func PositionOccupiedByPlayer(coordinate coordinate, goban *[19][19]position, player bool) bool {
	if CoordinateOnGoban(coordinate) == true {
		if PositionOccupied(coordinate, goban) == true {
			if SamePlayer(coordinate, goban, player) == true {
				return true
			}
		}
	}
	return false
}

func PositionOccupiedByOpponent(coordinate coordinate, goban *[19][19]position, player bool) bool {
	if CoordinateOnGoban(coordinate) == true {
		if PositionOccupied(coordinate, goban) == true {
			if SamePlayer(coordinate, goban, player) == false {
				return true
			}
		}
	}
	return false
}

func ThreeBlocked(end1 coordinate, end2 coordinate, goban *[19][19]position, player bool) bool {
	if PositionOccupiedByOpponent(end1, goban, player) == false && PositionOccupiedByOpponent(end2, goban, player) == false {
		return false
	}
	return true
}

func FindNeighbour(coordinate coordinate, y int8, x int8, multiple int8) coordinate {
	neighbour := coordinate
	neighbour.y += y * multiple
	neighbour.x += x * multiple
	return neighbour
}

func CheckVertexForThree(coordinate coordinate, goban *[19][19]position, y int8, x int8, player bool) (FreeThree bool) {
	minusTwo := FindNeighbour(coordinate, y, x, -2)
	minusOne := FindNeighbour(coordinate, y, x, -1)
	one := FindNeighbour(coordinate, y, x, 1)
	two := FindNeighbour(coordinate, y, x, 2)
	three := FindNeighbour(coordinate, y, x, 3)
	four := FindNeighbour(coordinate, y, x, 4)
	if PositionOccupiedByPlayer(one, goban, player) == true {
		if PositionOccupiedByPlayer(two, goban, player) == true {
			if ThreeBlocked(minusOne, three, goban, player) == false {
				return true
			}
		}
		if PositionOccupiedByPlayer(three, goban, player) == true {
			if ThreeBlocked(minusOne, four, goban, player) == false {
				if PositionOccupiedByOpponent(two, goban, player) == false {
					return true
				}
			}
		}
		if y < 0 || (y == 0 && x == -1) {
			if PositionOccupiedByPlayer(minusOne, goban, player) == true {
				if ThreeBlocked(minusTwo, two, goban, player) == false {
					return true
				}
			}
		}
	}
	if PositionOccupiedByPlayer(two, goban, player) == true {
		if PositionOccupiedByPlayer(three, goban, player) == true {
			if ThreeBlocked(minusOne, four, goban, player) == false {
				if PositionOccupiedByOpponent(one, goban, player) == false {
					return true
				}
			}
		}
		if PositionOccupiedByPlayer(minusOne, goban, player) == true {
			if ThreeBlocked(minusTwo, three, goban, player) == false {
				if PositionOccupiedByOpponent(one, goban, player) == false {
					return true
				}
			}
		}
	}
	return false
}

func DoubleThree(coordinate coordinate, g *game) (valid bool) { // returns false if move breaks rule
	var freeThree bool
	var x int8
	var y int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				foundThree := CheckVertexForThree(coordinate, &g.goban, y, x, g.player)
				if foundThree == true {
					if freeThree == true {
						return true
					} else {
						freeThree = true
					}
				}
			}
		}
	}
	return false
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

func ai(g *game) (coordinate coordinate, elapsed time.Duration) {
	start := time.Now()
	//	"Improved" Minimax implementation (Alpha-beta pruning, negascout, mtdf, ...) -> 5 points!
	//	Move search depth - 10 or more levels -> 5 points!
	//	search space of the algo - Multiple rectangular windows emcompassing placed stones but minimizing wasted space -> 5

	//	Heuristic
	//		take current alignments into account ?
	//		check whether an alignment has enough space to develop into a 5-in-a-row ?
	//		weigh an alignment according to its freedom (Free, half-free, flanked) ?
	//		take potential captures into account ?
	//		take current captured stones into account ?
	//		check for advanteageous combinations ?
	//		take both players into account ?
	//		take past player actions into account to identify patterns and weigh board states accordingly ?
	coordinate = RandomCoordinate()
	elapsed = time.Since(start)
	return coordinate, elapsed
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

func main() {
	fmt.Println("Hello world!") ////////
	g := InitializeGame()
	// launch ebiten. render goban and game stats 		// # do first
	// GameLoop(g)

	/// Test double three
	zero := coordinate{0, 0}  /////////
	PlaceIfValid(zero, g)     /////////
	three := coordinate{1, 1} /////////
	PlaceIfValid(three, g)    /////////
	three = coordinate{3, 5}  /////////
	PlaceIfValid(three, g)    /////////
	three = coordinate{3, 4}  /////////
	PlaceIfValid(three, g)    /////////
	g.player = true           //////
	three = coordinate{3, 2}  ///////// one of the three-aligned obstructed, therefore next move legal.
	PlaceIfValid(three, g)    /////////
	g.player = false          /////
	three = coordinate{3, 3}  ///////// Double Three, should be rejected
	PlaceIfValid(three, g)    ///////// Double Three, should be rejected

	/// Test Place stone
	// PlaceIfValid(RandomCoordinate(), g)
	// PlaceIfValid(RandomCoordinate(), g)
	// PlaceIfValid(RandomCoordinate(), g)
	// PlaceIfValid(RandomCoordinate(), g)
	// PlaceIfValid(RandomCoordinate(), g)
	// PlaceIfValid(RandomCoordinate(), g)
	// PlaceStone(RandomCoordinate(), true, &g.goban)  ////////
	// PlaceStone(RandomCoordinate(), true, &g.goban)  //////
	// PlaceStone(RandomCoordinate(), true, &g.goban)  ///////
	// PlaceStone(RandomCoordinate(), false, &g.goban) ////////
	// PlaceStone(RandomCoordinate(), true, &g.goban)  ////////

	// RemoveStone(zero, &g.goban) /////////

	coordinate, elapsed := ai(g)
	PlaceIfValid(coordinate, g) //////
	fmt.Println(elapsed)        // tme taken by ai //////////

	// gui.RunEbiten()

	DumpGoban(&g.goban)
	fmt.Println("Goodbye world!") ////////
}

// Rules:
// a stone, or solidly connected group of stones of one color, is captured if all liberties are occupied
// only allowed to place a stone in a position with no liberties if it immediately captures
