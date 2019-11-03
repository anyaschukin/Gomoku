package play

import (
	"math/rand"
)

func PlaceStone(coordinate coordinate, player bool, goban *[19][19]position) {
	goban[coordinate.y][coordinate.x].occupied = true
	goban[coordinate.y][coordinate.x].player = player
}

func IsMoveValid(coordinate coordinate, G *Game) bool {
	if PositionOccupied(coordinate, &G.goban) == true {
		G.message = "Position Occupied"
		return false
	}
	if DoubleThree(coordinate, G) == true {
		G.message = "Double-Three"
		return false
	}
	return true
}

func PlaceIfValid(coordinate coordinate, G *Game) bool { /// for human player
	if IsMoveValid(coordinate, G) == true {
		PlaceStone(coordinate, G.player, &G.goban)
		return true
	}
	return false
}

func RandomCoordinate() coordinate { ////////move this function somewhere else??
	x := int8(rand.Intn(19))
	y := int8(rand.Intn(19))
	random := coordinate{y, x}
	return random
}

func PlaceRandomIfValid(g *Game) (validated bool, coordinate coordinate) { //////// for testing
	coordinate = RandomCoordinate()
	if IsMoveValid(coordinate, G) == true {
		PlaceStone(coordinate, g.player, &g.goban)
		return true, coordinate
	}
	return false, coordinate
}
