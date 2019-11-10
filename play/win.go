package play

// capturedTen returns true if either Player has captured ten stones
func capturedTen(g *Game) (win bool) {
	if g.Capture0 >= 10 || g.Capture1 >= 10 {
		return true
	}
	return false
}

// checkVertexAlignFive returns true if 5 stones are aligned running through given coodinate on given axes
func checkVertexAlignFive(coordinate Coordinate, Goban *[19][19]position, y int8, x int8, Player bool) bool {
	var multiple int8
	var a int8
	var b int8
	for multiple = 1; multiple < 5; multiple++ {
		neighbour := FindNeighbour(coordinate, y, x, multiple)
		if PositionOccupiedByPlayer(neighbour, Goban, Player) == true {
			a++
		} else {
			break
		}
	}
	for multiple = -1; multiple > -5; multiple-- {
		neighbour := FindNeighbour(coordinate, y, x, multiple)
		if PositionOccupiedByPlayer(neighbour, Goban, Player) == true {
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
func AlignFive(coordinate Coordinate, Goban *[19][19]position, align5 *align5, Player bool, Capture0 uint8, Capture1 uint8) (alignedFive bool) {
	var x int8
	var y int8
	for y = -1; y <= 0; y++ {
		for x = -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				return false
			}
			if checkVertexAlignFive(coordinate, Goban, y, x, Player) == true {
				if canBreakFive(coordinate, Goban, y, x, Player) == true {
					align5.break5 = true
				}
				if canWinByCapture(Goban, Opponent(Player), Capture0, Capture1) == true {
					align5.capture8 = true
				}
				align5.winner = Player
				G.Winmove = coordinate
				return true
			}
		}
	}
	return false
}

func recordWin(G *Game, winner bool) {
	G.Won = true
	if winner == false {
		G.Message = "Black Wins!"
	} else {
		G.Message = "White Wins!"
	}
}

// CheckWin checks win conditions and updates Game struct
func CheckWin(coordinate Coordinate, G *Game) {
	if capturedTen(G) == true {
		recordWin(G, G.Player)
		G.Winmove = coordinate
		// fmt.Printf("Player %v wins by capturing 10.\n", G.Player)//////
	} else if G.align5.break5 == true {
		if PositionOccupiedByPlayer(G.Winmove, &G.Goban, G.align5.winner) == true &&
			AlignFive(G.Winmove, &G.Goban, &G.align5, G.align5.winner, G.Capture0, G.Capture1) == true {
			recordWin(G, Opponent(G.Player))
			// fmt.Printf("Player %v win by aligning 5.\nThe other Player could have broken this alignment by capturing a pair, but they didn't, silly!\nWinning move y:%d x:%d.\n", G.align5.winner, G.align5.winmove.y, G.align5.winmove.x)
		}
		G.align5.break5 = false
	} else if G.align5.capture8 == true {
		recordWin(G, Opponent(G.Player))
		// fmt.Printf("Player %v win by aligning 5.\nThe other Player could have Won by capturing ten, but they didn't, silly!\nWinning move y:%d x:%d.\n", G.align5.winner, G.align5.winmove.y, G.align5.winmove.x)
	}
	if AlignFive(coordinate, &G.Goban, &G.align5, G.Player, G.Capture0, G.Capture1) == true {
		if G.align5.break5 == false && G.align5.capture8 == false {
			recordWin(G, G.Player)
		}
		// if G.align5.break5 == true {
		// 	fmt.Printf("Player %v has aligned 5, however the other Player can break this alignment by capture\n", G.Player)
		// } else if G.align5.capture8 == true {
		// 	fmt.Printf("Player %v has aligned 5, however the other Player can win by capturing a pair\n", G.Player)
		// } else {
		// 	recordWin(G, G.Player)
		// 	fmt.Printf("Player %v wins by aligning 5! final move on position y:%d x:%d\n\n", G.Player, coordinate.y, coordinate.x)
		// }
	}
}
