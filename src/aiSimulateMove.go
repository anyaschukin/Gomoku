package gomoku

// ultimately, I'd like to re-use the same functions that Drew wrote originally
// that way, we're DRY

// re-write these functions so you're passing the goban and player, but not the whole game!

import "fmt"

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

// doubleThree returns true if suggested move breaks the double three rule
func doubleThree2(coordinate coordinate, goban *[19][19]position, player bool) bool {
	var freeThree bool
	var y int8
	var x int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				foundThree := checkVertexForThree(coordinate, goban, y, x, player)
				if foundThree == true {
					if freeThree == true {
						return true
					} else {
						freeThree = true
					}
				}
			}
		}
	}
	return false
}

func isMoveValid2(coordinate coordinate, goban *[19][19]position, player bool) bool {
	if coordinateOnGoban(coordinate) == false {
		return false
	}
	if positionOccupied(coordinate, goban) == true {
		return false
	}
	if doubleThree2(coordinate, goban, player) == true { // duplicate w/o *game
		return false
	}
	return true
}
