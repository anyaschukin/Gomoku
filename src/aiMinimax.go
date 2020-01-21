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
	// 	// fmt.Printf("child.id = %d, child.player = %v, child.maximizingPlayer: %v, child.coordinate: %v, child.value = %d\n", child.id, child.player, child.maximizingPlayer, child.coordinate, child.value)
	// }
	if maximizingPlayer == true {
		maxValue := minInt // set maxEval to -infinity
		for idx := range node.children {
			child := node.children[idx]
			maxValue = max(maxValue, child.value)
			value := minimaxRecursive(child, depth-1, alpha, beta, false)
			// fmt.Printf("\nvalue = %d, maxValue = %d \n", child.value, maxValue)
			alpha = max(alpha, value)
			if node.bestMove == nil || child.value == maxValue {
				node.bestMove = child
				// fmt.Printf("\nnode.bestMove.id = %d, node.bestMove.value = %d \n", node.bestMove.id, node.bestMove.value)
			}
			if beta <= alpha {
				break
			}
		}
		return maxValue
	} else {
		minValue := maxInt // set maxEval to +infinity
		for idx := range node.children {
			child := node.children[idx]
			minValue = min(minValue, child.value)
			value := minimaxRecursive(child, depth-1, alpha, beta, true)
			beta = min(beta, value)
			if node.bestMove == nil || child.value == minValue {
				node.bestMove = child
			}
			if beta <= alpha {
				break
			}
		}
		return minValue
	}
}
