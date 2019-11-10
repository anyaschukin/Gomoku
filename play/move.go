package play

import (
	"math/rand"
)

func PlaceStone(coordinate coordinate, Player bool, Goban *[19][19]position) {
	Goban[coordinate.y][coordinate.x].occupied = true
	Goban[coordinate.y][coordinate.x].Player = Player
}

func IsMoveValid(coordinate coordinate, G *Game) bool {
	if PositionOccupied(coordinate, &G.Goban) == true {
		G.message = "Position Occupied"
		return false
	}
	if DoubleThree(coordinate, G) == true {
		G.message = "Double-Three"
		return false
	}
	return true
}

func PlaceIfValid(coordinate coordinate, G *Game) bool { /// for human Player
	if IsMoveValid(coordinate, G) == true {
		PlaceStone(coordinate, G.Player, &G.Goban)
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
		PlaceStone(coordinate, g.Player, &g.Goban)
		return true, coordinate
	}
	return false, coordinate
}
