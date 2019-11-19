package play

import (
	"fmt"
	// "os"
	"time"
	// lib "Gomoku/golib"
	// "math"
)

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

var identity int

type node struct {
	id         int
	value      int
	goban      [19][19]position
	coordinate coordinate
	player     bool
	children   []*node
	bestMove   *node
}

func newNode(id int, value int, newGoban *[19][19]position, coordinate coordinate, newPlayer bool) *node {
	return &node{
		id:         id,
		value:      value, // change this to initialize to zero
		goban:      *newGoban,
		coordinate: coordinate,
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

// Recursively generates every move for a board (to depth 3), assigns value, and adds to tree
func generateBoardsDepth(depth int8, current *node, id int, player bool) {
	var y int8
	var x int8

	if depth > 2 {
		return
	}
	for y = 0; y < 19; y++ {
		for x = 0; x < 19; x++ {
			coordinate := coordinate{y, x}
			// fmt.Printf("depth %d\n", depth)
			if isMoveValid2(coordinate, &current.goban, player) == true { // duplicate of isMoveValid w/o *game
				identity++
				newGoban := current.goban
				placeStone(coordinate, player, &newGoban)

				value := valueBoard(&newGoban, player)
				child := newNode(identity, value, &newGoban, coordinate, player)
				addChild(current, current.id, child) //
				generateBoardsDepth(depth+1, child, child.id, !player)
			}
		}
	}
}

func createTree(g *game) *node {
	root := newNode(0, 0, &g.goban, coordinate{-1, -1}, g.player)
	generateBoardsDepth(1, root, root.id, root.player)
	return root
}

func printTree(parent *node) {
	current := parent
	fmt.Printf("\nparent: %d\n", current.id)
	for i := range current.children {
		child := current.children[i]
		fmt.Printf("child: %d", child.id)
		// printTree(child)
		// put in a mutex/lock to wait until this range is done, and then call printTree for the child
	}
	// depth-first
	for i := range current.children {
		current := current.children[i]
		// dumpGoban(&current.goban)
		printTree(current)
	}
}

func printBestRoute(root *node) {
	current := root
	for current.bestMove != nil {
		fmt.Println(current.id)
		fmt.Println(current.value)
		dumpGoban(&current.goban)
		current = current.bestMove
	}
	fmt.Println(current.id)
	fmt.Println(current.value)
	dumpGoban(&current.goban)
}

func minimaxTree(g *game) {
	start := time.Now()
	// depth := g.ai0.depth
	// if g.player == true {
	// 	depth = g.ai1.depth
	// }
	root := createTree(g)
	// printTree(root)
	// fmt.Println("-----")
	alpha := minInt
	beta := maxInt
	minimaxRecursive(root, 3, alpha, beta, true) // for some reason, maximizingplayer has to be set to 'false' for this to work
	elapsed := (time.Since(start))
	// printBestRoute(root)///
	// fmt.Printf("coordinate %v", root.bestMove.coordinate)

	fmt.Printf("Coordinate: %v --- player: %v\n", root.bestMove.coordinate, root.player)
	dumpGoban(&root.bestMove.goban)
	fmt.Println("------------\n")
	time.Sleep(100000000)

	if g.player == false {
		g.ai0.suggest = root.bestMove.coordinate
		g.ai0.timer = elapsed
	} else {
		g.ai1.suggest = root.bestMove.coordinate
		g.ai1.timer = elapsed
	}
}

//  creates a tree, whose root is the goban
//  creates all possible moves/boards to depth _, calculates values, add to tree
//  applies minimax to tree, finds best move

// addChild(root, 1, &node{id: 2, Value: 20})
// addChild(root, 1, &node{id: 3, Value: 30})
// addChild(root, 1, &node{id: 4, Value: 40})
// addChild(root, 2, &node{id: 5, Value: 50})
// addChild(root, 2, &node{id: 6, Value: 60})
// addChild(root, 2, &node{id: 7, Value: 70})
// addChild(root, 3, &node{id: 8, Value: 80})
// addChild(root, 3, &node{id: 9, Value: 90})
// addChild(root, 3, &node{id: 10, Value: 100})
// addChild(root, 4, &node{id: 11, Value: 110})
