package gomoku

import (
	"time"
)

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

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

//  Alpha is the tmp choice which has been found so far for the maximising player.
//  Beta is the tmp choice which has been found so far for the minimising player

func minimaxRecursive(node *node, depth uint8, alpha int, beta int, maximizingPlayer bool) int {

	if depth == 0 || node.value >= align5Win || node.value <= -align5Win {	
		return node.value
	}

	generateTree(node, node.coordinate, node.lastMove)

	if maximizingPlayer == true {
		maxValue := minInt // set value to -infinity
		for _, child := range node.children {
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
		for _, child := range node.children {
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

func minimaxTree(g *game) {
	start := time.Now()
	limit := g.ai0.depth
	if g.player == true {
		limit = g.ai1.depth
	}

	alpha := minInt
	beta := maxInt

	root := newNode(0, 0, &g.goban, g.lastMove, g.lastMove2, !g.player, false, g.capture0, g.capture1, nil, 1)
	minimaxRecursive(root, limit, alpha, beta, true)
	
	elapsed := (time.Since(start))

	if g.player == false {
		g.ai0.suggest = root.bestMove.coordinate
		g.ai0.timer = elapsed
	} else {
		g.ai1.suggest = root.bestMove.coordinate
		g.ai1.timer = elapsed
	}
}
