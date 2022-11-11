package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	daiba := os.Args[1:]
	var bas []int
	p := 96
	for _, a := range daiba {
		if a == "--upper" {
			p = 64
			continue
		}
		n := 0
		for _, j := range a {
			n = n*10 + int(rune(j)-'0')
		}
		bas = append(bas, n)
	}
	for i := 0; i < len(bas); i++ {
		if len(bas) == 0 {
			break
		} else if bas[i] > 26 {
			z01.PrintRune(rune(' '))
			continue
		}
		z01.PrintRune(rune(bas[i]) + rune(k))
	}
	if len(bas) > 0 {
		z01.PrintRune('\n')
	}
}
