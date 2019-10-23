package golib

import (
	"fmt"
	"os"
)

func Check(e error, message string) {
	if e != nil {
		fmt.Println(message)
		os.Exit(1)
	}
}
