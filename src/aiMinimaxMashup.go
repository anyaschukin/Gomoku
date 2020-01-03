package gomoku

// //  Alpha is the best choice which has been found so far for the maximising player.
// //  Beta is the best choice which has been found so far for the minimising player

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

func TreeMinimaxRecursive(node *node, depth uint8, alpha int, beta int, maximizingPlayer bool) int {

	if depth == 0 {
		return node.value
	}

	node.maximizingPlayer = maximizingPlayer
	generateBoardsDepth(node, node.id, node.player, node.coordinate, node.lastMove)
	if maximizingPlayer == true {
		maxValue := alpha // set maxEval to -infinity
		for idx := range node.children {
			child := node.children[idx]
			value := TreeMinimaxRecursive(child, depth-1, alpha, beta, false)
			maxValue = max(maxValue, value)
			alpha = max(alpha, value)
			if node.bestMove == nil || child.value > node.bestMove.value {
				node.bestMove = child
			}
			if beta <= alpha {
				break
			}
		}
		return maxValue
	} else {
		minValue := beta // set maxEval to +infinity
		for idx := range node.children {
			child := node.children[idx]
			value := TreeMinimaxRecursive(child, depth-1, alpha, beta, true)
			minValue = min(minValue, value)
			beta = min(beta, value)
			if node.bestMove == nil || child.value < node.bestMove.value {
				node.bestMove = child
			}
			if beta <= alpha {
				break
			}
		}
		return minValue
	}
}
