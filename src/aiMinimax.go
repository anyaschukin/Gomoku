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
		return node.value, node
	}

	/* DEBUG */
	// fmt.Printf("\nDEPTH = %d", depth)
	// fmt.Printf("\nparent.id = %d, parent.player = %v, parent.maximizingPlayer: %v, parent.coordinate: %v, parent.value = %d\n", node.id, node.player, node.maximizingPlayer, node.coordinate, node.value)
	// dumpGobanBlank(&node.goban)

	generateChildBoards(node, node.coordinate, node.lastMove)

	/* DEBUG */
	for i := range node.children {
	child := node.children[i]
	fmt.Printf("depth = %d, child.id = %d, child.player = %v, child.maximizingPlayer: %v, child.coordinate: %v, child.value = %d\n", depth, child.id, child.player, child.maximizingPlayer, child.coordinate, child.value)
	}

	var value int
	best := newNode(0, 0, &node.goban, node.coordinate, node.lastMove, !node.player, node.maximizingPlayer, node.captures.capture0, node.captures.capture1, node)
	bessst := newNode(0, 0, &node.goban, node.coordinate, node.lastMove, !node.player, node.maximizingPlayer, node.captures.capture0, node.captures.capture1, node)
	if maximizingPlayer == true {
		maxValue := minInt // set value to -infinity
		for idx := range node.children {
			child := node.children[idx]
			value, best = minimaxRecursive(child, depth-1, alpha, beta, false)
			fmt.Printf("value = %d, maxValue = %d\n", value, maxValue)
			// maxValue = max(value, maxValue)
			if value > maxValue {
				bessst = best
				maxValue = value
			}
			fmt.Printf("new maxValue = %d\n", maxValue)
			// if maxValue == best.value {
				// bessst = best
			// }
			alpha = max(alpha, maxValue)
			if alpha >= beta {
				break
			}
		}
		return maxValue, bessst
	} else {
		minValue := maxInt // set value to +infinity
		for idx := range node.children {
			child := node.children[idx]
			value, best = minimaxRecursive(child, depth-1, alpha, beta, true)
			fmt.Printf("value = %d, minValue = %d\n", value, minValue)
			// minValue = min(value, minValue)
			if value < minValue {
				bessst = best
				minValue = value
			}
			fmt.Printf("new minValue = %d\n", minValue)
			// if minValue == best.value {
				// bessst = best
			// }
			beta = min(beta, minValue)
			if beta <= alpha {
				break
			}
		}
		return minValue, bessst
	}
}
