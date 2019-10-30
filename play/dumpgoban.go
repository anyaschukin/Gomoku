package play

import (
	"fmt"
)

func DumpGoban(goban *[19][19]position) {
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

// fmt.Println(g.goban) // whole goban //////////
// fmt.Println(g.goban[0][0])          // one position ///////////
// fmt.Println(g.goban[0][0].occupied) // one position occupied /////////
// fmt.Println(g.goban[0][0].player)   // one position player  ///////////
