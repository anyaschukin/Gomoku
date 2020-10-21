package gomoku

//  Alpha is the tmp choice which has been found so far for the maximising player.
//  Beta is the tmp choice which has been found so far for the minimising player

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimaxRecursive(node *node, depth uint8, alpha int, beta int, maximizingPlayer bool) int {

	if depth == 0 || node.value >= align5Win {	
		return node.value
	}

	/* DEBUG */
	// fmt.Printf("\nDEPTH = %d", depth)
	// fmt.Printf("\nparent.id = %d, parent.player = %v, parent.maximizingPlayer: %v, parent.coordinate: %v, parent.value = %d\n", node.id, node.player, node.maximizingPlayer, node.coordinate, node.value)
	// dumpGobanBlank(&node.goban)

	generateChildBoards(node, node.coordinate, node.lastMove)

	/* DEBUG */
	// for i := range node.children {
	// child := node.children[i]
	// fmt.Printf("depth = %d, child.id = %d, child.player = %v, child.maximizingPlayer: %v, child.coordinate: %v, child.value = %d\n", depth, child.id, child.player, child.maximizingPlayer, child.coordinate, child.value)
	// }

	if maximizingPlayer == true {
		maxValue := minInt // set value to -infinity
		for idx := range node.children {
			child := node.children[idx]
			value := minimaxRecursive(child, depth-1, alpha, beta, false)
			if value > maxValue {
				node.bestMove = child
				maxValue = value
			}
			alpha = max(alpha, maxValue)
			if alpha >= beta {
				break
			}
		}
		return maxValue
	} else {
		minValue := maxInt // set value to +infinity
		for idx := range node.children {
			child := node.children[idx]
			value := minimaxRecursive(child, depth-1, alpha, beta, true)
			if value < minValue {
				node.bestMove = child
				minValue = value
			}
			beta = min(beta, minValue)
			if beta <= alpha {
				break
			}
		}
		return minValue
	}
}
