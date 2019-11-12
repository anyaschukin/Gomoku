package play

func removeStone(coordinate Coordinate, Goban *[19][19]position) {
	Goban[coordinate.Y][coordinate.X].occupied = false
}

func captureVertex(coordinate Coordinate, g *Game, y int8, x int8) {
	one := FindNeighbour(coordinate, y, x, 1)
	two := FindNeighbour(coordinate, y, x, 2)
	three := FindNeighbour(coordinate, y, x, 3)
	if PositionOccupiedByOpponent(one, &g.Goban, g.Player) == true &&
		PositionOccupiedByOpponent(two, &g.Goban, g.Player) == true &&
		PositionOccupiedByPlayer(three, &g.Goban, g.Player) == true {
		removeStone(one, &g.Goban)
		removeStone(two, &g.Goban)
		// fmt.Printf("Capture! Player: %v. captured y:%d x:%d & y:%d x:%d\n\n", g.Player, one.y, one.x, two.y, two.x) ///
		if g.Player == false {
			g.Capture0 += 2
		} else {
			g.Capture1 += 2
		}
	}
}

func Capture(coordinate Coordinate, g *Game) {
	var x int8
	var y int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				captureVertex(coordinate, g, y, x)
			}
		}
	}
}
