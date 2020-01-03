package gomoku

import (
	"fmt"

	// "os"
	"time"
	// lib "Gomoku/golib"
	// "math"
)

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

// const threatSpace = 3
var identity int

type node struct {
	id               int
	value            int
	goban            [19][19]position
	coordinate       coordinate
	lastMove         coordinate
	player           bool
	maximizingPlayer bool
	children         []*node
	bestMove         *node
}

func newNode(id int, value int, newGoban *[19][19]position, coordinate coordinate, lastMove coordinate, newPlayer bool) *node {
	return &node{
		id:         id,
		value:      value, // change this to initialize to zero
		goban:      *newGoban,
		coordinate: coordinate,
		lastMove:   lastMove,
		player:     newPlayer,
	}
}

// Recursively finds node by ID, and then appends child to node.chilren
func addChild(node *node, parentID int, child *node) int {
	if node.id == parentID {
		node.children = append(node.children, child)
		return 1
	} else {
		for idx, _ := range node.children {
			current := node.children[idx]
			addChild(current, parentID, child)
		}
	}
	return 0
}

func generateChildBoards(current *node, player bool, lastMove coordinate, x, y int8) {
	coordinate := coordinate{y, x}
	if isMoveValid2(coordinate, &current.goban, player) == true { // duplicate of isMoveValid w/o *game
		identity++
		newGoban := current.goban
		placeStone(coordinate, player, &newGoban)
		value := evaluateMove(coordinate, &newGoban, player)
		child := newNode(identity, value, &newGoban, coordinate, lastMove, player)
		addChild(current, current.id, child)
	}
}

// Generates every move for a board (to depth 3), assigns value, and adds to tree
func generateBoardsDepth(current *node, player bool, lastMove, lastMove2 coordinate) {
	var y int8
	var x int8

	for y = lastMove.y - 4; y <= lastMove.y+4; y++ {
		for x = lastMove.x - 4; x <= lastMove.x+4; x++ {
			generateChildBoards(current, player, lastMove, x, y)
			// coordinate := coordinate{y, x}
			// if isMoveValid2(coordinate, &current.goban, player) == true { // duplicate of isMoveValid w/o *game
			// 	identity++
			// 	newGoban := current.goban
			// 	placeStone(coordinate, player, &newGoban)
			// 	value := evaluateMove(coordinate, &newGoban, player)
			// 	child := newNode(identity, value, &newGoban, coordinate, lastMove, player)
			// 	addChild(current, current.id, child)
			// }
		}
	}
	for y = lastMove2.y - 4; y <= lastMove2.y+4; y++ {
		for x = lastMove2.x - 4; x <= lastMove2.x+4; x++ {
			// optimized so the threat-space searches don't overlap
			if !(y >= lastMove.y-4 && y <= lastMove.y+4 && x >= lastMove.x-4 && x <= lastMove.x+4) {
				generateChildBoards(current, player, lastMove2, x, y)
				// coordinate := coordinate{y, x}
				// if isMoveValid2(coordinate, &current.goban, player) == true { // duplicate of isMoveValid w/o *game
				// 	identity++
				// 	newGoban := current.goban
				// 	placeStone(coordinate, player, &newGoban)
				// 	value := evaluateMove(coordinate, &newGoban, player)
				// 	child := newNode(identity, value, &newGoban, coordinate, lastMove2, player)
				// 	addChild(current, current.id, child)
				// }
			}
		}
	}
}

/* prints the tree from the root */
func printTree(parent *node) {
	current := parent
	fmt.Printf("\nparent: %d\n", current.id)
	for i := range current.children {
		child := current.children[i]
		fmt.Printf("child: %d", child.id) //////
		// printTree(child)
		// put in a mutex/lock to wait until this range is done, and then call printTree for the child
	}
	/* depth-first */
	for i := range current.children {
		current := current.children[i]
		// dumpGoban(&current.goban)
		printTree(current)
	}
}

/* prints the best move at the selected depth */
func printBestRoute(root *node) {
	current := root
	for current.bestMove != nil {
		fmt.Printf("id = %d, value = %d, maximizingPlayer = %v\n", current.id, current.value, current.maximizingPlayer)
		// dumpGoban(&current.goban)
		current = current.bestMove
	}
	fmt.Printf("id = %d, value = %d, maximizingPlayer = %v\n\n", current.id, current.value, current.maximizingPlayer)
	// dumpGoban(&current.goban)
}

func minimaxTree(g *game) {
	start := time.Now()
	limit := g.ai0.depth
	if g.player == true {
		limit = g.ai1.depth
	}

	root := newNode(0, 0, &g.goban, g.lastMove, g.lastMove2, g.player)
	alpha := minInt
	beta := maxInt
	TreeMinimaxRecursive(root, limit, alpha, beta, true)
	elapsed := (time.Since(start))
	// printBestRoute(root)
	// fmt.Printf("Coordinate: %v , eval: %v , player: %v\n", root.bestMove.coordinate, root.bestMove.value, root.player)
	// dumpGoban(&root.bestMove.goban)

	if g.player == false {
		g.ai0.suggest = root.bestMove.coordinate
		g.ai0.timer = elapsed
	} else {
		g.ai1.suggest = root.bestMove.coordinate
		g.ai1.timer = elapsed
	}
}
