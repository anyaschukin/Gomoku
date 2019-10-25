package main

import (
	"fmt"
	"math/rand"
	"time"
	//	lib "Gomoku/golib"
	//	gui "Gomoku/GUI"
)

type coordinate struct {
	x int8
	y int8
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
	g.player = true ///// rm, just to test
	return &g
}

func RandomCoordinate() coordinate {
	x := int8(rand.Intn(19))
	y := int8(rand.Intn(19))
	random := coordinate{x, y}
	return random
}

func PlaceStone(coordinate coordinate, player bool, goban *[19][19]position) {
	goban[coordinate.x][coordinate.y].occupied = true
	goban[coordinate.x][coordinate.y].player = player
}

func RemoveStone(coordinate coordinate, goban *[19][19]position) {
	goban[coordinate.x][coordinate.y].occupied = false
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

func PositionOccupied(coordinate coordinate, goban *[19][19]position) (occupied bool) {
	if goban[coordinate.x][coordinate.y].occupied == true {
		return true
	}
	return false
}

func SamePlayer(coordinate coordinate, goban *[19][19]position, player bool) (samePlayer bool) {
	if goban[coordinate.x][coordinate.y].player == player {
		return true
	}
	return false
}

func IsCoordinateOnGoban(coordinate coordinate) (onGoban bool) {
	if coordinate.x < 0 || coordinate.x > 18 || coordinate.y < 0 || coordinate.y > 18 {
		return false
	}
	return true
}

func CheckVertexForThree(coordinate coordinate, goban *[19][19]position, x int8, y int8, player bool) (FreeThree bool) {
	one := coordinate
	one.x += x
	one.y += y
	two := coordinate
	two.x += x * 2
	two.y += y * 2
	if IsCoordinateOnGoban(one) == true {
		if PositionOccupied(one, goban) == true && SamePlayer(one, goban, player) == true {
			if IsCoordinateOnGoban(two) == true {
				if PositionOccupied(two, goban) == true && SamePlayer(two, goban, player) == true {
					return true
				}
			}
			three := coordinate
			three.x += x * 3
			three.y += y * 3
			if IsCoordinateOnGoban(three) == true {
				if PositionOccupied(three, goban) == true && SamePlayer(three, goban, player) == true {
					return true
				}
			}
			if y < 0 || (y == 0 && x == -1) {
				minusOne := coordinate
				minusOne.x -= x
				minusOne.y -= y
				if IsCoordinateOnGoban(minusOne) == true {
					if PositionOccupied(minusOne, goban) == true && SamePlayer(minusOne, goban, player) == true {
						return true
					}
				}
			}
		}
	}
	if IsCoordinateOnGoban(two) == true {
		if PositionOccupied(two, goban) == true && SamePlayer(two, goban, player) == true {
			three := coordinate
			three.x += x * 3
			three.y += y * 3
			if IsCoordinateOnGoban(three) == true {
				if PositionOccupied(three, goban) == true && SamePlayer(three, goban, player) == true {
					return true
				}
			}
			minusOne := coordinate
			minusOne.x -= x
			minusOne.y -= y
			if IsCoordinateOnGoban(minusOne) == true {
				if PositionOccupied(minusOne, goban) == true && SamePlayer(minusOne, goban, player) == true {
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
				foundThree := CheckVertexForThree(coordinate, &g.goban, x, y, g.player)
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
	if g.goban[coordinate.x][coordinate.y].occupied == true {
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

	fmt.Println(g.goban) // whole goban //////////
	// fmt.Println(g.goban[0][0])          // one position ///////////
	// fmt.Println(g.goban[0][0].occupied) // one position occupied /////////
	// fmt.Println(g.goban[0][0].player)   // one position player  ///////////
	fmt.Println("Goodbye world!") ////////
}

// Rules:
// a stone, or solidly connected group of stones of one color, is captured if all liberties are occupied
// only allowed to place a stone in a position with no liberties if it immediately captures
