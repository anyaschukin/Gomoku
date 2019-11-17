package ai

// import (
// 	"fmt"

// 	//	lib "Gomoku/golib"

// 	play "Gomoku/play"
// )

//  create tree
//  generate board/node
//  assign value to board
//  plug board into minimax
//  choose board

func GameLoop(G *play.Game) {
	validated := false
	coordinate := play.RandomCoordinate() /////
	for i := 0; i < 10000; i++ {          //moves
		validated, coordinate = play.PlaceRandomIfValid(G)
		if validated == true {
			play.Capture(coordinate, G)
			play.DumpGoban(&G.Goban)   //////
			play.CountStones(&G.Goban) /////////
			play.CheckWin(coordinate, G)
			G.Player = play.SwapPlayers(G.Player)
		}
	}
	// update game.moves ++
	//	return err
}

func main() {
	fmt.Println("Hello world!") ////////
	G := play.newGame()
	// gui.RunEbiten()
	GameLoop(G)
	fmt.Println("Goodbye world!") ////////
}

// func GameLoop(g) {
// if AI:git ad
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
// PlaceStone(coordinate, true, &g.Goban)
// check if capture
//		remove captured
//		update game.captured struct
// check win
//		5 in a row
// 		all win conditions?
// update game. struct
//		Player change
//		moves ++
// re-render ebiten with updated Goban and stats
//	return err
// }
