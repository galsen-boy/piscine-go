package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lennbr := 0
	bil := false
	li := 96
	// argument := os.Args
	for in, correspond := range os.Args {
		lennbr = in
		if correspond == "---upper" {
			bil = true
		}
	}
	if bil == true {
		li = 64
	}
	for i := 1; i <= lennbr; i++ {
		for _, correspond := range os.Args[i] {
			z01.PrintRune(correspond + rune(li))
		}
		z01.PrintRune(10)
	}
}
