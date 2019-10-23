package main

import (
	"fmt"
)

// type coordinate struct{
//     x		uint8
//     y		uint8
// }

type position struct{// alternatively uint8 (0 = unocupied), but memory waste
    occupied	bool
    player		bool
}

type game struct {
    goban		[19][19]position
    player		bool				// who's move is it? (player 0 - black first)
//	capture0	uint8				// capture 10 and win
//	capture1	uint8				// capture 10 and win
//	move		uint32				// how many moves have been played in total (is this desirable/necessary?)
//	pass		bool				// was the last move a pass (if next move pass game over)
//	last		coordinate/*position?			// last move to check ko rule ()
}

type config struct { 	/// merge with game struct?
//	aiplayer	bool	// is player 1 human or AI
//	hotseat		bool	// AI player only suggests moves, human must choose move
//	prescience	uint8	// how many moves in advance do we examine
}

func InitializeGame() *game{
	g := game{}
	return &g
}

//func PlaceStone(coordinate, player, *goban) {
//	g.goban[coordinate.x][coordinate.y].occupied = True // untested
//	g.goban[coordinate.x][coordinate.y].player = player // untested
//}

// func GameLoop(g) {
	// if AI:
	//		suggest move
	// if human or hotseat:
	//		listen for mouse click 						//###### do first
	//			find position/pass request clicked 		//###### do first
	//			if pass, double pass end?
	//			if reset, reset game with new config
	// 		check if position is valid (if human, assume ai has aleady checked)
	//			occupied?
	//			rules
	//				ko
	//				double-three
	// place stone(coordinate, player, *goban) 									//###### do first
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
// }

func main() {
	fmt.Println("Hello world!")////////
	g := InitializeGame()
	// launch ebiten. render goban and game stats 		// # do first
	// GameLoop(g)

	fmt.Println(g.goban)// whole goban //////////
	fmt.Println(g.goban[0][0])// one position ///////////
	fmt.Println(g.goban[0][0].occupied)// one position occupied /////////
	fmt.Println(g.goban[0][0].player)// one position player  ///////////
	fmt.Println("Goodbye world!")////////
}
