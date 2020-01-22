package gomoku

// import "fmt"

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

func minimaxRecursive(node *node, depth uint8, alpha int, beta int, maximizingPlayer bool) int {

	if depth == 0 {
		return node.value
	}

	// fmt.Printf("\nparent.id = %d, parent.player = %v, parent.maximizingPlayer: %v, parent.coordinate: %v, parent.value = %d\n", node.id, node.player, node.maximizingPlayer, node.coordinate, node.value)
	generateBoards(node, node.coordinate, node.lastMove)

	// for i := range node.children {
	// 	child := node.children[i]
	// 	fmt.Printf("child.id = %d, child.player = %v, child.maximizingPlayer: %v, child.coordinate: %v, child.value = %d\n", child.id, child.player, child.maximizingPlayer, child.coordinate, child.value)
	// }
	if maximizingPlayer == true {
		maxValue := minInt // set maxEval to -infinity
		for idx := range node.children {
			child := node.children[idx]
			value := minimaxRecursive(child, depth-1, alpha, beta, false)
			maxValue = max(maxValue, value)
			alpha = max(alpha, value)
			if node.id == 0 && value == maxValue {
				node.bestMove = child
			}
			// if node.bestMove == nil || value == maxValue {
				// node.bestMove = child
			// }
			if beta <= alpha {
				break
			}
		}
		return maxValue
	} else {
		minValue := maxInt // set maxEval to +infinity
		for idx := range node.children {
			child := node.children[idx]
			value := minimaxRecursive(child, depth-1, alpha, beta, true)
			minValue = min(minValue, value)
			beta = min(beta, value)
			if node.id == 0 && value == minValue {
				node.bestMove = child
			}
			// if node.bestMove == nil || value == minValue {
				// node.bestMove = child
			// }
			if beta <= alpha {
				break
			}
		}
		return minValue
	}
}
