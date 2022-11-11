package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	compteur := 0
	for s := range argument {
		compteur = s + 1
	}
	for a := 1; a < compteur; a++ {
		for b := a + 1; b < compteur; b++ {
			if argument[a] > argument[b] {
				argument[a], argument[b] = argument[b], argument[a]
			}
		}
	}
	for b := 1; b <= compteur-1; b++ {
		for _, daiba := range argument[b] {
			z01.PrintRune(daiba)
		}
		z01.PrintRune('\n')
	}
}
