package ai

import (
	"fmt"
	"math"
	// play "Gomoku.play"
)

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

type Node struct {
	ID    int
	Value int
	// game     *play.Game
	Children []*Node
	BestMove *Node
}

func NewNode(ID int, value int /*, Game *play.Game*/) *Node {
	return &Node{
		ID:    ID,
		Value: value,
	}
	// node.board = &play.Game
	// return Node
}

func AddChild(node *Node, parentID int, child *Node) int {
	if node.ID == parentID {
		node.Children = append(node.Children, child)
		return 1
	} else {
		for idx, _ := range node.Children {
			current := node.Children[idx]
			AddChild(current, parentID, child)
		}
	}
	return 0
}

func createTree() *Node {
	root := NewNode(1, 10)

	AddChild(root, 1, &Node{ID: 2, Value: 20})
	AddChild(root, 1, &Node{ID: 3, Value: 30})
	AddChild(root, 1, &Node{ID: 4, Value: 40})
	AddChild(root, 2, &Node{ID: 5, Value: 50})
	AddChild(root, 2, &Node{ID: 6, Value: 60})
	AddChild(root, 2, &Node{ID: 7, Value: 70})
	AddChild(root, 3, &Node{ID: 8, Value: 80})
	AddChild(root, 3, &Node{ID: 9, Value: 90})
	AddChild(root, 3, &Node{ID: 10, Value: 100})
	AddChild(root, 4, &Node{ID: 11, Value: 110})

	return root
}

func printTree(parent *Node) {
	current := parent
	fmt.Printf("\nparent: %d\n", current.ID)
	for i := range current.Children {
		child := current.Children[i]
		fmt.Printf("child: %d", child.ID)
		// put in a mutex/lock to wait until this range is done, and then call printTree for the child
	}
	for i := range current.Children {
		current := current.Children[i]
		printTree(current)
	}
}

func Tree() {
	root := createTree()
	// printTree(root)
	fmt.Println("-----")
	// alpha := float64(math.Inf(-1))
	// beta := float64(math.Inf(1))
	alpha := MinInt
	beta := MaxInt
	MinimaxRecursive(root, 3, alpha, beta, false)
	current := root
	for current.BestMove != nil {
		fmt.Println(current.ID)
		current = current.BestMove
	}
	fmt.Println(current.ID)
}