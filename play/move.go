package play

import (
	"fmt"
	"math/rand"
)

func PlaceStone(coordinate coordinate, player bool, goban *[19][19]position) {
	goban[coordinate.y][coordinate.x].occupied = true
	goban[coordinate.y][coordinate.x].player = player
}

func IsMoveValid(coordinate coordinate, g *game) (whyInvalid string) {
	if PositionOccupied(coordinate, &g.goban) == true {
		return "Position already Occupied"
	}
	if DoubleThree(coordinate, g) == true {
		return "Move introduces a forbidden double-three"
	}
	//	other rules?? ko?
	return "Valid"
}

func PlaceIfValid(coordinate coordinate, g *game) { /// for human player
	whyInvalid := IsMoveValid(coordinate, g)
	if whyInvalid == "Valid" {
		PlaceStone(coordinate, g.player, &g.goban)
	} else {
		// return whyInvalid to gui
		fmt.Println(whyInvalid) /////
	}
}

func RandomCoordinate() coordinate { ////////move this function somewhere else??
	x := int8(rand.Intn(19))
	y := int8(rand.Intn(19))
	random := coordinate{y, x}
	return random
}

func PlaceRandomIfValid(g *game) (validated bool, coordinate coordinate) { //////// for testing
	coordinate = RandomCoordinate()
	whyInvalid := IsMoveValid(coordinate, g)
	if whyInvalid == "Valid" {
		PlaceStone(coordinate, g.player, &g.goban)
		return true, coordinate
	} else {
		// return whyInvalid to gui
		fmt.Println(whyInvalid) /////
		return false, coordinate
	}
}
