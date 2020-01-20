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

func minimaxRecursive(node *node, depth uint8, startDepth uint8, alpha int, beta int, maximizingPlayer bool) int {

	if depth == 0 {
		return node.value
	}

	generateBoards(node, node.coordinate, node.lastMove)

	fmt.Printf("parent.id = %d, parent.player = %v, parent.maximingPlayer = %v, parent.coordinate: %v, parent.value = %d\n", node.id, node.player, node.maximizingPlayer, node.coordinate, node.value) //////
	// if node.id == 357130 {
	// 	dumpGobanBlank(&node.goban)
	// }
	for i := range node.children {
		child := node.children[i]
		fmt.Printf("child.id = %d, child.player = %v, child.maximizingPlayer: %v, child.coordinate: %v, child.value = %d\n", child.id, child.player, child.maximizingPlayer, child.coordinate, child.value) //////
		// if child.value >= 2039951306 {
		// 	fmt.Printf("Oh Hi\n") //////
		// }
		// if child.id == 361550 {
		// 	dumpGobanBlank(&node.goban)
		// }
	}
	fmt.Printf("\n") //////

	if maximizingPlayer == true {
		maxValue := minInt // set maxEval to -infinity
		for idx := range node.children {
			child := node.children[idx]
			value := minimaxRecursive(child, depth-1, startDepth, alpha, beta, false)
			// fmt.Printf("id = %d, player = %v, maximizingPlayer: %v, coordinate: %v, value = %d\n", child.id, child.player, child.maximizingPlayer, child.coordinate, child.value) //////
			// fmt.Printf("maxValue: %v, value: %v\n", maxValue, value)
			maxValue = max(maxValue, value)
			// alpha = max(alpha, value)
			// if depth == startDepth && (cooronate not set or child.value > best move value) {
				// best move coordinate = child.coordinate
			//}
			// if node.bestMove == nil || child.value > node.bestMove.value {
			if node.bestMove == nil || value == maxValue {
				node.bestMove = child
			}
			// if beta <= alpha {
			// 	break
			// }
		}
		// if depth == startDepth {
		// 	return node.bestMove.id /// best move coordinate
		// } else {
			fmt.Printf("maxValue: %v\n\n", maxValue)
			return maxValue
		// }
	} else {
		minValue := maxInt // set maxEval to +infinity
		for idx := range node.children {
			child := node.children[idx]
			value := minimaxRecursive(child, depth-1, startDepth, alpha, beta, true)
			// fmt.Printf("id = %d, player = %v, maximizingPlayer: %v, coordinate: %v, value = %d\n", child.id, child.player, child.maximizingPlayer, child.coordinate, child.value) //////
			// fmt.Printf("minValue: %v, value: %v\n", minValue, value)
			minValue = min(minValue, value)
			// beta = min(beta, value)
			// if node.bestMove == nil || child.value < node.bestMove.value {
			if node.bestMove == nil || value == minValue {
				node.bestMove = child
			}
			// if beta <= alpha {
			// 	break
			// }
		}
		fmt.Printf("minValue: %v\n\n", minValue)
		return minValue
	}
}
