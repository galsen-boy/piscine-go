package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args

	var taille int
	for l := range argument {
		taille = l
	}

	for i := taille; i > 0; i-- {
		for _, j := range argument[i] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
