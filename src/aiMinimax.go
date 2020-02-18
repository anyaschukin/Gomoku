package gomoku

import "fmt"

//  Alpha is the best choice which has been found so far for the maximising player.
//  Beta is the best choice which has been found so far for the minimising player

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

func minimaxRecursive(node *node, depth uint8, alpha int, beta int, maximizingPlayer bool) (int, *node) {

	if depth == 0 {
		// fmt.Printf("WIN node.id = %d, node.coordinate = %v, node.value = %d\n", node.id, node.coordinate, node.value)
		return node.value, node /*node.id, node.coordinate*/
	}

	/* DEBUG */
	fmt.Printf("\nDEPTH = %d", depth)
	fmt.Printf("\nparent.id = %d, parent.player = %v, parent.maximizingPlayer: %v, parent.coordinate: %v, parent.value = %d\n", node.id, node.player, node.maximizingPlayer, node.coordinate, node.value)
	// dumpGobanBlank(&node.goban)

	generateChildBoards(node, node.coordinate, node.lastMove)

	/* DEBUG */
	for i := range node.children {
		child := node.children[i]
		fmt.Printf("child.id = %d, child.player = %v, child.maximizingPlayer: %v, child.coordinate: %v, child.value = %d\n", child.id, child.player, child.maximizingPlayer, child.coordinate, child.value)
	}

	var value int
	best := newNode(0, 0, &node.goban, node.coordinate, node.lastMove, !node.player, node.maximizingPlayer, node.captures.capture0, node.captures.capture1, node)
	bessst := newNode(0, 0, &node.goban, node.coordinate, node.lastMove, !node.player, node.maximizingPlayer, node.captures.capture0, node.captures.capture1, node)
	if maximizingPlayer == true {
		// value := minInt // set value to -infinity
		maxValue := minInt // set value to -infinity
		for idx := range node.children {
			child := node.children[idx]
			// value = max(value, minimaxRecursive(child, depth-1, alpha, beta, false))
			value, best = minimaxRecursive(child, depth-1, alpha, beta, false)
			maxValue = max(value, maxValue)
			if maxValue == best.value {
				bessst = best
			}
			// fmt.Printf("value = %d\n", value)
			alpha = max(alpha, value)
			// if node.id == 0 && child.value == maxValue {
			// 	id = child.id
			// }
			// if node.id == 0 && (node.bestMove == nil || node.value == value) {
			// node.bestMove = child
			// }
			// if node.bestMove == nil || value > node.bestMove.value {
			// node.bestMove = child
			// }
			if node.id == 0 && (node.bestMove == nil || child.value > node.bestMove.value) {
				node.bestMove = child
			}
			if beta <= alpha {
				break
			}
		}
		return maxValue, bessst
	} else {
		// value := maxInt // set value to +infinity
		minValue := maxInt // set value to +infinity
		// var id int
		for idx := range node.children {
			child := node.children[idx]
			// value, id = min(value, minimaxRecursive(child, depth-1, alpha, beta, true))
			value, best = minimaxRecursive(child, depth-1, alpha, beta, true)
			minValue = min(value, minValue)
			beta = min(beta, value)
			if minValue == best.value {
				bessst = best
			}
			// if node.id == 0 /*&& value == minValue*/ {
			// node.bestMove = child
			// }
			if node.bestMove == nil || child.value < node.bestMove.value {
				node.bestMove = child
			}
			if beta <= alpha {
				break
			}
		}
		return minValue, bessst
	}
}
