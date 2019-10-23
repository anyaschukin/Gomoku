package main

import (
	"fmt"
)

type position struct{// alternatively uint8 (0 = unocupied), but memory waste
    occupied	bool
    player		bool
}

type game struct {
    goban		[19][19]position
    player		bool// who's move is it? (player 0 - black first)
//	capture0	uint8 //capture 10 and win
//	capture1	uint8 //capture 10 and win
}

type config struct { /// merge with game struct?
//	player0ai	bool	// is player 0 human or ai
//	player1ai	bool	// is player 0 human or ai
//	prescience	uint8	// how many moves in advance do we examine
}

func InitializeGame() *game{
	g := game{}
	return &g
}

func main() {
	fmt.Println("Hello world!")////////
	g := InitializeGame()
	fmt.Println(g.goban)// whole goban //////////
	fmt.Println(g.goban[0][0])// one position ///////////
	fmt.Println(g.goban[0][0].occupied)// one position occupied /////////
	fmt.Println(g.goban[0][0].player)// one position player  ///////////
	fmt.Println("Goodbye world!")////////
}
