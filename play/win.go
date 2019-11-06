package play

// capturedTen returns true if either player has captured ten stones
func capturedTen(g *Game) (win bool) {
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
		neighbour := FindNeighbour(coordinate, y, x, multiple)
		if PositionOccupiedByPlayer(neighbour, goban, player) == true {
			a++
		} else {
			break
		}
	}
	for multiple = -1; multiple > -5; multiple-- {
		neighbour := FindNeighbour(coordinate, y, x, multiple)
		if PositionOccupiedByPlayer(neighbour, goban, player) == true {
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

// AlignFive returns true if 5 stones are aligned running through given coodinate
func AlignFive(coordinate coordinate, goban *[19][19]position, align5 *align5, player bool, capture0 uint8, capture1 uint8) (alignedFive bool) {
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
				if canWinByCapture(goban, Opponent(player), capture0, capture1) == true {
					align5.capture8 = true
				}
				align5.winner = player
				G.winmove = coordinate
				return true
			}
		}
	}
	return false
}

func recordWin(G *Game, winner bool) {
	G.won = true
	if winner == false {
		G.message = "Black Wins!"
	} else {
		G.message = "White Wins!"
	}
}

// CheckWin checks win conditions and updates Game struct
func CheckWin(coordinate coordinate, G *Game) {
	if capturedTen(G) == true {
		recordWin(G, G.player)
		G.winmove = coordinate
		// fmt.Printf("Player %v wins by capturing 10.\n", G.player)//////
	} else if G.align5.break5 == true {
		if PositionOccupiedByPlayer(G.winmove, &G.goban, G.align5.winner) == true &&
			AlignFive(G.winmove, &G.goban, &G.align5, G.align5.winner, G.capture0, G.capture1) == true {
			recordWin(G, Opponent(G.player))
			// fmt.Printf("Player %v win by aligning 5.\nThe other player could have broken this alignment by capturing a pair, but they didn't, silly!\nWinning move y:%d x:%d.\n", G.align5.winner, G.align5.winmove.y, G.align5.winmove.x)
		}
		G.align5.break5 = false
	} else if G.align5.capture8 == true {
		recordWin(G, Opponent(G.player))
		// fmt.Printf("Player %v win by aligning 5.\nThe other player could have won by capturing ten, but they didn't, silly!\nWinning move y:%d x:%d.\n", G.align5.winner, G.align5.winmove.y, G.align5.winmove.x)
	}
	if AlignFive(coordinate, &G.goban, &G.align5, G.player, G.capture0, G.capture1) == true {
		if G.align5.break5 == false && G.align5.capture8 == false {
			recordWin(G, G.player)
		}
		// if G.align5.break5 == true {
		// 	fmt.Printf("Player %v has aligned 5, however the other player can break this alignment by capture\n", G.player)
		// } else if G.align5.capture8 == true {
		// 	fmt.Printf("Player %v has aligned 5, however the other player can win by capturing a pair\n", G.player)
		// } else {
		// 	recordWin(G, G.Player)
		// 	fmt.Printf("Player %v wins by aligning 5! final move on position y:%d x:%d\n\n", G.player, coordinate.y, coordinate.x)
		// }
	}
}
