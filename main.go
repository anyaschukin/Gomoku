package main

import (
	// "fmt"

	//	lib "Gomoku/golib"

	play "Gomoku/play"
)

// func GameLoop(g) {
// if AI:
//		suggest move
// if human or hotseat:
//		listen for mouse click 						//###### do first
//			find position/pass/new/exit clicked 		//###### do first
//			if pass, double pass end?
//			if reset, reset Game with new config
// 		check if position is valid (if human, assume ai has aleady checked)
//			occupied?
//			rules
//				ko
//				double-three
// PlaceStone(coordinate, true, &g.goban)
// check if capture
//		remove captured
//		update Game.captured struct
// check win
//		5 in a row
// 		all win conditions?
// update Game. struct
//		player change
//		moves ++
// re-render ebiten with updated goban and stats
//	return err
// }

func main() {
	// fmt.Println("Hello world!") ////////
	play.Play()
	// fmt.Println("Goodbye world!") ////////
}
