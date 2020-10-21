package gomoku

import (
	"fmt"
)

// * Some minimax tree debugging funcs * //

// Optional add: 
// dumpGoban16(&current.goban)

// Prints the tree from the root
func printTree(parent *node) {
	// fmt.Printf("\n---------TREE-----------\n")
	current := parent
	fmt.Printf("\nparent: %d\n", current.id)
	for i := range current.children {
		child := current.children[i]
		fmt.Printf("child: %d", child.id)
		printTree(child)
	}
}

// Finds the parent node from any given child or grandchild node
func findParent(leaf *node) *node {
	fmt.Printf("\n\n----------findParent----------\n")
	current := leaf
	fmt.Printf("id = %d, coordinate = %v, lastMove = %d, value = %d, player = %v, maximizingPlayer = %v, parent.id = %d\n", current.id, current.coordinate, current.lastMove, current.value, current.player, current.maximizingPlayer, current.parent.id)//////////////!!!!!!!!
	for current.parent.id != 0 {
		current = current.parent
		fmt.Printf("id = %d, coordinate = %v, lastMove = %d, value = %d, player = %v, maximizingPlayer = %v, parent.id = %d\n", current.id, current.coordinate, current.lastMove, current.value, current.player, current.maximizingPlayer, current.parent.id)//////////////!!!!!!!!
	}
	fmt.Printf("\n")
	return current
}