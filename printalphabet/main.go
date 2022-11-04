package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	a := []rune(name)
	for i := range a {
		z01.PrintRune(a[i])
	}
	z01.PrintRune(10)
}
