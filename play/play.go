package play

// import (
// 	"fmt"
// 	// lib "Gomoku/golib"
// 	// gui "Gomoku/GUI"
// )



type coordinate struct {
	y int8
	x int8
}

type position struct {
	occupied bool
	player   bool
}

type align5 struct { //winning move for checking if opponent breaks it in the next move
	break5   bool
	capture8 bool // is it possible for the opponent to win by capturing 10? (have they already captured 8, and is there an available capture move)
	winner   bool /// rm?
	winmove  coordinate
}

type Game struct {
	goban    [19][19]position
	player   bool   // whose move is it? (player 0 - black first)
	capture0 uint8  // capture 10 and win
	capture1 uint8  // capture 10 and win
	align5   align5 // one player has aligned 5, however it can be broken. The other player must break it, capture 10, or lose.
	// move		uint32				// how many moves have been played in total (is this desirable/necessary?)
	// input      *Input
	// boardImage *ebiten.Image ///
}

// type mouseState int

// const (
// 	mouseStateNone mouseState = iota
// 	mouseStatePressing
// 	mouseStateSettled
// )

// type Input struct {
// 	mouseState    mouseState
// 	mousePosX int
// 	mousePosY int
// }


var (
	G *Game
)

// func init() {
// 	G := &Game{}
// }

// type ai struct { 	/// merge with Game struct?
//	aiplayer	bool	// is player 1 human or AI
//	hotseat		bool	// AI player only suggests moves, human must choose move
//	prescience	uint8	// how many moves in advance do we examine
// }

// func InitializeGame() *Game {
// 	G := Game{}
// 	// g.player = true ///// rm, just to test
// 	return &G
// }

func NewGame() *Game {
	G = &Game{}
	return G
}

func SwapPlayers(player bool) bool {
	if player == false {
		return true
	} else {
		return false
	}
}

func GameLoop(G *Game) {//(G *Game) {
	// G = Ga
	validated := false
	coordinate := RandomCoordinate() /////
	for i := 0; i < 10; i++ {       //moves ////!!!!!!
		validated, coordinate = PlaceRandomIfValid(G)
		if validated == true {
			Capture(coordinate, G)
			DumpGoban(&G.goban) //////
			// CountStones(&G.goban) /////////
			CheckWin(coordinate, G)
			G.player = SwapPlayers(G.player)
		}
	}
	// update Game.moves ++
	//	return err
	// return G
}

func Play() {
	G := NewGame()
	G.player = false
	RunEbiten()
}

// func Play() {
// G := initializeGame()
// GameLoop(G)

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

// // launch ebiten. render goban and Game stats 		// # do first
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
// }
