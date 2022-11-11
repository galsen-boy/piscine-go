package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lennbr := 0
	bool := false
	li := 96
	// argument := os.Args
	for i, correspond := range os.Args {
		lennbr = index
		if correspond == "---upper" {
			bool = true
		}
	}
	if bool == true {
		li = 64
	}
	for i := 1; i <= lennbr; i++ {
		for _, correspond := range os.Args[i] {
			z01.PrintRune(correspond + rune(li))
		}
		z01.PrintRune(10)
	}
}
