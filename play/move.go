package play

import (
	"fmt"
	"math/rand"
)

func PlaceStone(coordinate coordinate, player bool, goban *[19][19]position) {
	goban[coordinate.y][coordinate.x].occupied = true
	goban[coordinate.y][coordinate.x].player = player
}

func RemoveStone(coordinate coordinate, goban *[19][19]position) {
	goban[coordinate.y][coordinate.x].occupied = false
}

func RandomCoordinate() coordinate {
	x := int8(rand.Intn(19))
	y := int8(rand.Intn(19))
	random := coordinate{y, x}
	return random
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
