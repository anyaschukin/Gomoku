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
	player           bool // black or white
	maximizingPlayer bool // used by miniMax algo
	children         []*node
	bestMove         *node
}

func newNode(id int, value int, newGoban *[19][19]position, coordinate coordinate, lastMove coordinate, newPlayer bool, maximizingPlayer bool) *node {
	return &node{
		id:               id,
		value:            value, // change this to initialize to zero
		goban:            *newGoban,
		coordinate:       coordinate, // change this to move
		lastMove:         lastMove,
		player:           newPlayer,
		maximizingPlayer: maximizingPlayer,
	}
}

// Recursively finds node by ID, and then appends child to node.chilren
func addChild(node *node, parentID int, child *node) {
	if node.id == parentID {
		node.children = append(node.children, child)
	} else {
		for idx, _ := range node.children {
			current := node.children[idx]
			addChild(current, parentID, child)
		}
	}
}

// Generates every move for a board, assigns value, and adds to tree
func generateChildBoards(current *node, lastMove coordinate, x, y int8) {
	var value int
	coordinate := coordinate{y, x}
	if isMoveValid2(coordinate, &current.goban, current.player) == true { // duplicate of isMoveValid w/o *game
		identity++
		newGoban := current.goban
		placeStone(coordinate, !current.player, &newGoban)
		if current.maximizingPlayer == true {
			value = current.value - evaluateMove(coordinate, &newGoban, current.player)
		} else {
			value = current.value + evaluateMove(coordinate, &newGoban, current.player)
		}
		child := newNode(identity, value, &newGoban, coordinate, lastMove, !current.player, !current.maximizingPlayer)
		addChild(current, current.id, child)
	}
}

func generateBoards(current *node, lastMove, lastMove2 coordinate) {
	var y int8
	var x int8

	// fmt.Printf("current.player: %v\n", current.player)
	for y = lastMove.y - 4; y <= lastMove.y+4; y++ {
		for x = lastMove.x - 4; x <= lastMove.x+4; x++ {
			generateChildBoards(current, lastMove, x, y)
		}
	}
	for y = lastMove2.y - 4; y <= lastMove2.y+4; y++ {
		for x = lastMove2.x - 4; x <= lastMove2.x+4; x++ {
			// optimized so the threat-space searches don't overlap
			if !(y >= lastMove.y-4 && y <= lastMove.y+4 && x >= lastMove.x-4 && x <= lastMove.x+4) {
				generateChildBoards(current, lastMove2, x, y)
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
		dumpGoban(&current.goban)
		printTree(current)
	}
}

/* prints the best move at the selected depth */
func printBestRoute(root *node) {
	current := root
	fmt.Printf("root.player = %v\n", root.player)
	for current.bestMove != nil {
		fmt.Printf("id = %d, value = %d, move = %v, maximizingPlayer = %v\n", current.id, current.value, current.coordinate, current.maximizingPlayer)
		// dumpGoban(&current.goban)//////// standard out
		// dumpGobanBlank(&current.goban) //////// file
		current = current.bestMove
	}
	fmt.Printf("id = %d, value = %d, move = %v, maximizingPlayer = %v\n\n", current.id, current.value, current.coordinate, current.maximizingPlayer)
	// dumpGoban(&current.goban)//////// standard out
	// dumpGobanBlank(&current.goban) ///////// file
}

func minimaxTree(g *game) {
	start := time.Now()
	limit := g.ai0.depth
	if g.player == true {
		limit = g.ai1.depth
	}

	root := newNode(0, 0, &g.goban, g.lastMove, g.lastMove2, !g.player, false)
	// fmt.Printf("First root.player = %v\n", root.player)
	alpha := minInt
	beta := maxInt
	value_wtf := minimaxRecursive(root, limit, limit, alpha, beta, true)
	// minimaxRecursive(root, limit, alpha, beta, true)
	fmt.Printf("value_wtf: %v\n\n", value_wtf) //////////
	elapsed := (time.Since(start))
	printBestRoute(root)                                                 /////////////
	fmt.Printf("\n\n----------------------------------------------\n\n") //////////
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
