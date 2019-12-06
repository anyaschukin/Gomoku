package gomoku

// capturedTen returns true if either Player has captured ten stones
func capturedTen(g *game) bool {
	if g.capture0 >= 10 || g.capture1 >= 10 {
		return true
	}
	return false
}

// measureChain returns how many stones in a row for given coordinate, axes & player
func measureChain(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) int8 {
	var length int8
	var multiple int8
	for multiple = 1; multiple < 5; multiple++ {
		neighbour := findNeighbour(coordinate, y, x, multiple)
		if positionOccupiedByPlayer(neighbour, goban, player) == true {
			length++
		} else {
			break
		}
	}
	return length
}

// checkVertexAlignFive returns true if 5 stones are aligned running through given coodinate on given axes
func checkVertexAlignFive(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) bool {
	a := measureChain(coordinate, goban, y, x, player)
	b := measureChain(coordinate, goban, -y, -x, player)
	if a+b >= 4 {
		return true
	}
	return false
}

// alignFive returns true if 5 stones are aligned running through given coodinate
func alignFive(coordinate coordinate, goban *[19][19]position, align5 *align5, player bool, capture0, capture1 uint8) bool {
	if positionOccupiedByPlayer(coordinate, &g.goban, player) == false {
		return false
	}
	var x int8
	var y int8
	for y = -1; y <= 0; y++ {
		for x = -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				return false
			}
			if checkVertexAlignFive(coordinate, goban, y, x, player) == true {
				recordAlignFive(coordinate, y, x, g)
				if canBreakFive(coordinate, goban, y, x, player) == true {
					align5.break5 = true
				}
				if canWinByCapture(goban, opponent(player), capture0, capture1) == true {
					align5.capture8 = true
				}
				return true
			}
		}
	}
	return false
}

// opponent returns the opponent of the current Player
func opponent(player bool) bool {
	if player == false {
		return true
	}
	return false
}

func recordWin(g *game, winner bool) {
	g.won = true
	if winner == false {
		g.gui.message = "Black Wins!"
	} else {
		g.gui.message = "White Wins!"
	}
}

// checkWin checks win conditions and updates Game struct
func checkWin(coordinate coordinate, g *game) {
	if capturedTen(g) == true {
		recordWin(g, g.player)
		g.winMove = coordinate
	} else if g.align5.capture8 == true {
		// Opponent wins by aligning 5. Player could have won by capturing ten, but didn't, silly!
		recordWin(g, opponent(g.player))
	} else if g.align5.break5 == true {
		if positionOccupiedByPlayer(g.winMove, &g.goban, opponent(g.player)) == true &&
			alignFive(g.winMove, &g.goban, &g.align5, opponent(g.player), g.capture0, g.capture1) == true {
			// Opponent wins by aligning 5. Player could have broken this alignment by capturing, but didn't, silly!
			recordWin(g, opponent(g.player))
			return
		}
		g.align5.break5 = false
	}
	if alignFive(coordinate, &g.goban, &g.align5, g.player, g.capture0, g.capture1) == true {
		if g.align5.break5 == false && g.align5.capture8 == false {
			// Player wins by aligning 5!
			recordWin(g, g.player)
		}
	}
}
