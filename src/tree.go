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

type node struct {
	id		int
	value   int
	goban   [19][19]position 
	player	bool
	// game     *play.Game
	children []*node
	bestMove *node
}

// newGoban := [19][19]position{}

func newNode(id int, value int, newGoban *[19][19]position, newPlayer bool) *node {
	return &node{
		id:   	id,
		value:	value,	// change this to initialize to zero
		goban:	*newGoban,
		player: newPlayer,
	}
	// node.board = &play.Game
	// return node
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
			if isMoveValid2(coordinate, &current.goban, player) == true {		// duplicate of isMoveValid w/o *game
				// id not incrementing correctly in recursive function
				id += 1
				newGoban := current.goban
				placeStone(coordinate, player, &newGoban)
				value := valueBoard(&newGoban, player)
				child := newNode(id, value, &newGoban, player)
				addChild(current, current.id, child) //
				generateBoardsDepth(depth+1, child, child.id, !player)
			}
			// continue
		}
	}
}

func createTree(g *game) *node {
	id := 1
	root := newNode(id, 10, &g.goban, g.player)
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
	time.Sleep(1000000000)
	// depth-first
	for i := range current.children {
		current := current.children[i]
		// dumpGoban(&current.goban)
		printTree(current)
	}
}

func minimaxTree(g *game) {
	root := createTree(g)
	printTree(root)
	// fmt.Println("-----")
	// alpha := float64(math.Inf(-1))
	// beta := float64(math.Inf(1))
	// alpha := minInt
	// beta := maxInt
	// minimaxRecursive(root, 3, alpha, beta, false)	// for some reason, maximizingplayer has to be set to 'false' for this to work
	// current := root
	// for current.bestMove != nil {
		// fmt.Println(current.id)
		// current = current.bestMove
	// }
	// fmt.Println(current.id)
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