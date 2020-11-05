package gomoku

import (
	"fmt"
)

// dumpGoban prints the Goban with color
// original dumpGoban that works on 42 Macs
func dumpGoban(goban *[19][19]position) {
	fmt.Printf("     ")
	for x := 0; x < 19; x++ {
		fmt.Printf("{%2d         } ", x)
	}
	fmt.Printf("\n")
	for y := 0; y < 19; y++ {
		fmt.Printf("{%2d} ", y)
		for x := 0; x < 19; x++ {
			if goban[y][x].occupied == true {
				fmt.Printf("\x1B[31m")
			}
			fmt.Printf("{%v\x1B[0m ", goban[y][x].occupied)
			if goban[y][x].occupied == true {
				fmt.Printf(" ")
			}
			color := ""
			if goban[y][x].occupied == true {
				if goban[y][x].player == true {
					color = "\x1B[32m"
				} else {
					color = "\x1B[33m"
				}
			}
			fmt.Printf("%s%v\x1B[0m", color, goban[y][x].player)
			if goban[y][x].player == true {
				fmt.Printf(" ")
			}
			fmt.Printf("} ")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// dumpGoban prints the Goban with color
// specifically for Anya and Drew's 16in laptops
func dumpGoban16(goban *[19][19]position) {
	fmt.Printf("     ")
	for x := 0; x < 19; x++ {
		fmt.Printf("{%2d         }", x)
	}
	fmt.Printf("\n")
	for y := 0; y < 19; y++ {
		fmt.Printf("{%2d} ", y)
		for x := 0; x < 19; x++ {
			if goban[y][x].occupied == true {
				fmt.Printf("\x1B[31m")
			}
			fmt.Printf("{%v\x1B[0m ", goban[y][x].occupied)
			if goban[y][x].occupied == true {
				fmt.Printf(" ")
			}
			color := ""
			if goban[y][x].occupied == true {
				if goban[y][x].player == true {
					color = "\x1B[32m"
				} else {
					color = "\x1B[33m"
				}
			}
			fmt.Printf("%s%v\x1B[0m", color, goban[y][x].player)
			if goban[y][x].player == true {
				fmt.Printf(" ")
			}
			fmt.Printf("}")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// dumpGobanBlank prints the Goban without color 
// cleaner output when printing to file
func dumpGobanBlank(goban *[19][19]position) {
	for x := 0; x < 19; x++ {
		fmt.Printf("{%2d        } ", x)
	}
	fmt.Printf("\n")
	for y := 0; y < 19; y++ {
		for x := 0; x < 19; x++ {
			fmt.Printf("%v ", goban[y][x].occupied)
			if goban[y][x].occupied == true {
				fmt.Printf(" ")
			}
			fmt.Printf("%v", goban[y][x].player)
			if goban[y][x].player == true {
				fmt.Printf(" ")
			}
			fmt.Printf("} ")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}


// * Some minimax tree debugging funcs * //

// Prints the tree from the root
func printTree(parent *node) {
	fmt.Printf("\n---------TREE-----------\n")
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
