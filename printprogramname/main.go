package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	daiba := []rune(os.Args[0])
	for i := range daiba {
		if i > 1 {
			z01.PrintRune(daiba[i])
		}
	}
	z01.PrintRune('\n')
}
