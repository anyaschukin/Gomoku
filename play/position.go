package play

func CoordinateOnGoban(coordinate coordinate) (onGoban bool) {
	if coordinate.y < 0 || coordinate.y > 18 || coordinate.x < 0 || coordinate.x > 18 {
		return false
	}
	return true
}

func PositionOccupied(coordinate coordinate, Goban *[19][19]position) (occupied bool) {
	if Goban[coordinate.y][coordinate.x].occupied == true {
		return true
	}
	return false
}

func SamePlayer(coordinate coordinate, Goban *[19][19]position, Player bool) (samePlayer bool) {
	if Goban[coordinate.y][coordinate.x].Player == Player {
		return true
	}
	return false
}

func PositionOccupiedByPlayer(coordinate coordinate, Goban *[19][19]position, Player bool) bool {
	if CoordinateOnGoban(coordinate) == true {
		if PositionOccupied(coordinate, Goban) == true &&
			SamePlayer(coordinate, Goban, Player) == true {
			return true
		}
	}
	return false
}

func PositionOccupiedByOpponent(coordinate coordinate, Goban *[19][19]position, Player bool) bool {
	if CoordinateOnGoban(coordinate) == true {
		if PositionOccupied(coordinate, Goban) == true &&
			SamePlayer(coordinate, Goban, Player) == false {
			return true
		}
	}
	return false
}

func PositionUnoccupied(coordinate coordinate, Goban *[19][19]position) (unoccupied bool) {
	if CoordinateOnGoban(coordinate) == true {
		if PositionOccupied(coordinate, Goban) == false {
			return true
		}
	}
	return false
}

func FindNeighbour(coordinate coordinate, y int8, x int8, multiple int8) coordinate {
	neighbour := coordinate
	neighbour.y += y * multiple
	neighbour.x += x * multiple
	return neighbour
}
