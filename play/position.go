package play

func coordinateOnGoban(coordinate Coordinate) (onGoban bool) {
	if coordinate.Y < 0 || coordinate.Y > 18 || coordinate.X < 0 || coordinate.X > 18 {
		return false
	}
	return true
}

func PositionOccupied(coordinate Coordinate, Goban *[19][19]position) (occupied bool) {
	if Goban[coordinate.Y][coordinate.X].occupied == true {
		return true
	}
	return false
}

func SamePlayer(coordinate Coordinate, Goban *[19][19]position, Player bool) (samePlayer bool) {
	if Goban[coordinate.Y][coordinate.X].Player == Player {
		return true
	}
	return false
}

func PositionOccupiedByPlayer(coordinate Coordinate, Goban *[19][19]position, Player bool) bool {
	if coordinateOnGoban(coordinate) == true {
		if PositionOccupied(coordinate, Goban) == true &&
			SamePlayer(coordinate, Goban, Player) == true {
			return true
		}
	}
	return false
}

func PositionOccupiedByOpponent(coordinate Coordinate, Goban *[19][19]position, Player bool) bool {
	if coordinateOnGoban(coordinate) == true {
		if PositionOccupied(coordinate, Goban) == true &&
			SamePlayer(coordinate, Goban, Player) == false {
			return true
		}
	}
	return false
}

func PositionUnoccupied(coordinate Coordinate, Goban *[19][19]position) (unoccupied bool) {
	if coordinateOnGoban(coordinate) == true {
		if PositionOccupied(coordinate, Goban) == false {
			return true
		}
	}
	return false
}

func FindNeighbour(coordinate Coordinate, y int8, x int8, multiple int8) Coordinate {
	neighbour := coordinate
	neighbour.Y += y * multiple
	neighbour.X += x * multiple
	return neighbour
}
