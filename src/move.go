package play

// import (
// 	"math/rand"
// )

func PlaceStone(coordinate Coordinate, Player bool, Goban *[19][19]position) {
	Goban[coordinate.Y][coordinate.X].occupied = true
	Goban[coordinate.Y][coordinate.X].Player = Player
}

func IsMoveValid(coordinate Coordinate, G *Game) bool {
	if PositionOccupied(coordinate, &G.Goban) == true {
		G.Message = "Position Occupied"
		return false
	}
	if DoubleThree(coordinate, G) == true {
		G.Message = "Double-Three"
		return false
	}
	return true
}

func PlaceIfValid(coordinate Coordinate, G *Game) bool { /// for human Player
	if IsMoveValid(coordinate, G) == true {
		PlaceStone(coordinate, G.Player, &G.Goban)
		return true
	}
	return false
}

// func Randomcoordinate() Coordinate { ////////move this function somewhere else??
// 	x := int8(rand.Intn(19))
// 	y := int8(rand.Intn(19))
// 	random := Coordinate{y, x}
// 	return random
// }

// func PlaceRandomIfValid(g *Game) (validated bool, coordinate Coordinate) { //////// for testing
// 	coordinate = Randomcoordinate()
// 	if IsMoveValid(coordinate, G) == true {
// 		PlaceStone(coordinate, g.Player, &g.Goban)
// 		return true, coordinate
// 	}
// 	return false, coordinate
// }
