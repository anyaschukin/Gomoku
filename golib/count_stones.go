package golib

// import "fmt"/////

// func CountStones(goban *[19][19]position) (black uint16, white uint16) {
// 	var y int8
// 	var x int8
// 	player := false // black
// 	for y = 0; y < 19; y++ {
// 		for x = 0; x < 19; x++ {
// 			coordinate := coordinate{y, x}
// 			if positionOccupiedByPlayer(coordinate, goban, player) == true {
// 				black++
// 			}
// 			if positionOccupiedByOpponent(coordinate, goban, player) == true {
// 				white++
// 			}
// 		}
// 	}
// 	// fmt.Printf("Count stones, black: %d, white: %d\n", black, white) ////////
// 	return
// }
