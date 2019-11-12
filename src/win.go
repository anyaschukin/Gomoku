package play

// capturedTen returns true if either Player has captured ten stones
func capturedTen(g *game) (win bool) {
	if g.capture0 >= 10 || g.capture1 >= 10 {
		return true
	}
	return false
}

// checkVertexAlignFive returns true if 5 stones are aligned running through given coodinate on given axes
func checkVertexAlignFive(coordinate coordinate, goban *[19][19]position, y int8, x int8, player bool) bool {
	var multiple int8
	var a int8
	var b int8
	for multiple = 1; multiple < 5; multiple++ {
		neighbour := findNeighbour(coordinate, y, x, multiple)
		if positionOccupiedByPlayer(neighbour, goban, player) == true {
			a++
		} else {
			break
		}
	}
	for multiple = -1; multiple > -5; multiple-- {
		neighbour := findNeighbour(coordinate, y, x, multiple)
		if positionOccupiedByPlayer(neighbour, goban, player) == true {
			b++
		} else {
			break
		}
	}
	if a+b >= 4 {
		return true
	}
	return false
}

// alignFive returns true if 5 stones are aligned running through given coodinate
func alignFive(coordinate coordinate, goban *[19][19]position, align5 *align5, player bool, capture0 uint8, capture1 uint8) (alignedFive bool) {
	var x int8
	var y int8
	for y = -1; y <= 0; y++ {
		for x = -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				return false
			}
			if checkVertexAlignFive(coordinate, goban, y, x, player) == true {
				if canBreakFive(coordinate, goban, y, x, player) == true {
					align5.break5 = true
				}
				if canWinBycapture(goban, opponent(player), capture0, capture1) == true {
					align5.capture8 = true
				}
				align5.winner = player
				g.Winmove = coordinate
				return true
			}
		}
	}
	return false
}

func recordWin(g *game, winner bool) {
	g.Won = true
	if winner == false {
		g.Message = "Black Wins!"
	} else {
		g.Message = "White Wins!"
	}
}

// checkWin checks win conditions and updates Game struct
func checkWin(coordinate coordinate, g *game) {
	if capturedTen(g) == true {
		recordWin(g, g.player)
		g.Winmove = coordinate
		// fmt.Printf("Player %v wins by capturing 10.\n", g.Player)//////
	} else if g.align5.break5 == true {
		if positionOccupiedByPlayer(g.Winmove, &g.goban, g.align5.winner) == true &&
			alignFive(g.Winmove, &g.goban, &g.align5, g.align5.winner, g.capture0, g.capture1) == true {
			recordWin(g, opponent(g.player))
			// fmt.Printf("Player %v win by aligning 5.\nThe other Player could have broken this alignment by capturing a pair, but they didn't, silly!\nWinning move y:%d x:%d.\n", g.align5.winner, g.align5.winmove.y, g.align5.winmove.x)
		}
		g.align5.break5 = false
	} else if g.align5.capture8 == true {
		recordWin(g, opponent(g.player))
		// fmt.Printf("Player %v win by aligning 5.\nThe other Player could have Won by capturing ten, but they didn't, silly!\nWinning move y:%d x:%d.\n", g.align5.winner, g.align5.winmove.y, g.align5.winmove.x)
	}
	if alignFive(coordinate, &g.goban, &g.align5, g.player, g.capture0, g.capture1) == true {
		if g.align5.break5 == false && g.align5.capture8 == false {
			recordWin(g, g.player)
		}
		// if g.align5.break5 == true {
		// 	fmt.Printf("Player %v has aligned 5, however the other Player can break this alignment by capture\n", g.Player)
		// } else if g.align5.capture8 == true {
		// 	fmt.Printf("Player %v has aligned 5, however the other Player can win by capturing a pair\n", g.Player)
		// } else {
		// 	recordWin(G, g.Player)
		// 	fmt.Printf("Player %v wins by aligning 5! final move on position y:%d x:%d\n\n", g.Player, coordinate.y, coordinate.x)
		// }
	}
}
