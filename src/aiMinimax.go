package gomoku

import "fmt"

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

	fmt.Printf("\nDEPTH = %d", depth)
	fmt.Printf("\nparent.id = %d, parent.player = %v, parent.maximizingPlayer: %v, parent.coordinate: %v, parent.value = %d\n", node.id, node.player, node.maximizingPlayer, node.coordinate, node.value)
	generateBoards(node, node.coordinate, node.lastMove)

	for i := range node.children {
		child := node.children[i]
		fmt.Printf("child.id = %d, child.player = %v, child.maximizingPlayer: %v, child.coordinate: %v, child.value = %d\n", child.id, child.player, child.maximizingPlayer, child.coordinate, child.value)
	}
	if maximizingPlayer == true {
		// maxValue := minInt // set maxEval to -infinity
		value := minInt // set maxEval to -infinity
		for idx := range node.children {
			child := node.children[idx]
			value = max(value, minimaxRecursive(child, depth-1, alpha, beta, false))
			fmt.Printf("\nmaxValue = %d\n", value)
			// maxValue = max(maxValue, value)
			alpha = max(alpha, value)
			// if node.id == 0 && value == maxValue {
			// node.bestMove = child
			// }
			if node.bestMove == nil || child.value > node.bestMove.value {
				if node.bestMove != nil {
					fmt.Printf("Max - child.id = %d, child.value = %d, node.bestMove.value = %d\n\n", child.id, child.value, node.bestMove.value)
				}
				node.bestMove = child
			}
			if beta <= alpha {
				break
			}
		}
		// return maxValue
		return value
	} else {
		// minValue := maxInt // set maxEval to +infinity
		value := maxInt // set maxEval to +infinity
		for idx := range node.children {
			child := node.children[idx]
			value = min(value, minimaxRecursive(child, depth-1, alpha, beta, true))
			fmt.Printf("\nminValue = %d\n", value)
			// minValue = min(minValue, value)
			beta = min(beta, value)
			// if node.id == 0 && value == minValue {
			// node.bestMove = child
			// }
			if node.bestMove == nil || child.value < node.bestMove.value {
				if node.bestMove != nil {
					fmt.Printf("Min - child.id = %d, child.value = %d, node.bestMove.value = %d\n\n", child.id, child.value, node.bestMove.value)
				}
				node.bestMove = child
			}
			if beta <= alpha {
				break
			}
		}
		return value
	}
}
