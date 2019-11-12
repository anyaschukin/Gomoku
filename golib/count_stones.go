package play

// import "fmt"/////

func CountStones(Goban *[19][19]position) (black uint16, white uint16) {
	var y int8
	var x int8
	player := false // black
	for y = 0; y < 19; y++ {
		for x = 0; x < 19; x++ {
			coordinate := coordinate{y, x}
			if positionOccupiedByPlayer(coordinate, Goban, player) == true {
				black++
			}
			if positionOccupiedByOpponent(coordinate, Goban, player) == true {
				white++
			}
		}
	}
	// fmt.Printf("Count stones, black: %d, white: %d\n", black, white) ////////
	return
}
