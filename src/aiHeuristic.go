package gomoku

// chainAttackDefend returns a score for aligning or blocking a chain of 5, 4, 3, or 2 stones
func chainAttackDefend(coordinate coordinate, goban *[19][19]position, y, x int8, player bool) int {
	var attack int
	var defend int

	opponentChain, flanked1, flanked2 := lengthOpponentChain(coordinate, goban, y, x, player)
	switch opponentChain {
	case 4:
		defend = blockWin
	case 3:
		if flanked1 == false && flanked2 == false {
			defend = block3Free
		}
	case 2:
		if flanked1 == false && flanked2 == false {
			defend = block2
		}
	}

	playerChain, flanked1, flanked2 := lengthPlayerChain(coordinate, goban, y, x, player)
	switch playerChain {
	case 5:
		attack = align5Win
	case 4:
		if flanked1 == false && flanked2 == false {
			attack = align4Free
		} else if flanked1 == true && flanked2 == true {
			attack = 0
		} else {
			attack = align4FLanked
		}
	case 3:
		if flanked1 == false && flanked2 == false {
			attack = align3Free
		} else if flanked1 == true && flanked2 == true {
			attack = 0
		} else {
			attack = align3Flanked
		}
	case 2:
		if flanked1 == false && flanked2 == false {
			attack = align2Free
		}
	}
	return attack + defend
}

// evaluateMove checks for alignments/captures along each vertex for one move, and returns a score for that move
func evaluateMove(coordinate coordinate, goban *[19][19]position, player bool, captures captures) int {
	var x int8
	var y int8

	eval := 0
	for y = -1; y <= 0; y++ {
		for x = -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				return eval
			}
			capt := captureAttackDefend(coordinate, goban, y, x, player, captures)
			if capt >= blockWin || capt <= -blockWin {
				return capt
			}
			eval += capt
			tmp := chainAttackDefend(coordinate, goban, y, x, player)
			if tmp >= blockWin || tmp <= -blockWin {
				return tmp
			}
			eval += tmp
		}
	}
	return eval
}
