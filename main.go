package main

import (
	gomoku "Gomoku/src"
	// "fmt"
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()
	// for i := 0; i <= 60000; i++ {
	// 	fmt.Printf("WHY U NO WORK\n")
	// }
	gomoku.Play()
}

/// go run main.go
