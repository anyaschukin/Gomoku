package play

import (
	"fmt"
)

func DumpGoban(Goban *[19][19]position) {
	fmt.Printf("     ")
	for x := 0; x < 19; x++ {
		fmt.Printf("{%2d         } ", x)
	}
	fmt.Printf("\n")
	for y := 0; y < 19; y++ {
		fmt.Printf("{%2d} ", y)
		for x := 0; x < 19; x++ {
			if Goban[y][x].occupied == true {
				fmt.Printf("\x1B[31m")
			}
			fmt.Printf("{%v\x1B[0m ", Goban[y][x].occupied)
			if Goban[y][x].occupied == true {
				fmt.Printf(" ")
			}
			color := ""
			if Goban[y][x].occupied == true {
				if Goban[y][x].player == true {
					color = "\x1B[32m"
				} else {
					color = "\x1B[33m"
				}
			}
			fmt.Printf("%s%v\x1B[0m", color, Goban[y][x].player)
			if Goban[y][x].player == true {
				fmt.Printf(" ")
			}
			fmt.Printf("} ")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// fmt.Println(g.Goban) // whole Goban //////////
// fmt.Println(g.Goban[0][0])          // one position ///////////
// fmt.Println(g.Goban[0][0].occupied) // one position occupied /////////
// fmt.Println(g.Goban[0][0].Player)   // one position Player  ///////////
