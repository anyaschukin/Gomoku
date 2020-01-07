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

	node.maximizingPlayer = maximizingPlayer
	generateBoards(node, node.coordinate, node.lastMove)
	for i := range node.children {
		child := node.children[i]
		fmt.Printf("child.id = %d, child.value = %d\n", child.id, child.value) //////
	}
	fmt.Printf("\n") //////
	// fmt.Printf("node id: %v\n", node.id)
	// fmt.Printf("node value: %v\n", node.value)
	// fmt.Printf("node coordinate: %v\n", node.coordinate)
	// fmt.Printf("node id: %v\n", node.id)
	// fmt.Printf("node id: %v\n", node.id)
	// fmt.Printf("node id: %v\n", node.id)
	// fmt.Printf("node id: %v\n", node.id)
	// fmt.Printf("node id: %v\n", node.id)
	// fmt.Printf("node id: %v\n", node.id)
	// id               int
	// value            int
	// goban            [19][19]position
	// coordinate       coordinate
	// lastMove         coordinate
	// player           bool // black or white
	// maximizingPlayer bool // used by miniMax algo
	// children         []*node
	// bestMove         *node

	if maximizingPlayer == true {
		maxValue := alpha // set maxEval to -infinity
		for idx := range node.children {
			child := node.children[idx]
			value := minimaxRecursive(child, depth-1, alpha, beta, false)
			// if len(node.children) == 0 {
			// 	value = node.value
			// }
			maxValue = max(maxValue, value)
			alpha = max(alpha, value)
			// if node.bestMove == nil || child.value == maxValue {
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
			value := minimaxRecursive(child, depth-1, alpha, beta, true)

			minValue = min(minValue, value)
			beta = min(beta, value)
			// if node.bestMove == nil || child.value == minValue {
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
