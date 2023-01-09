package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args

	for i := range argument {
		if i != 0 {
			for a := range argument[i] {
				b := []rune(argument[i])
				{
					z01.PrintRune(b[a])
				}
			}
			z01.PrintRune('\n')
		}
	}
}
