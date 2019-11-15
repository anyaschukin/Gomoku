package ai

import "math"

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

func minimaxRecursive(node *node, depth int, alpha int, beta int, maximizingPlayer bool) int {
	// if game over in position {
	if depth == 0 || len(node.Children) == 0 {
		return node.value
	}

	if maximizingPlayer {
		maxValue := alpha // set maxEval to -infinity
		// generate new boards (by creating a board with a piece at the next spot)
		for idx, _ := range node.Children {
			child := node.Children[idx]
			value := minimaxRecursive(child, depth-1, alpha, beta, false)
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
		for idx, _ := range node.Children {
			child := node.Children[idx]
			value := minimaxRecursive(child, depth-1, alpha, beta, true)
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
