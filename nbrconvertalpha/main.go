package main

import (
	"os"

	"github.com/01-edu/z01"
)

func BasicAtoi(s string) int {
	x := 0
	for _, n := range s {
		y := 0
		for i := '1'; i <= n; i++ {
			y++
		}
		x = x*10 + y
	}
	return x
}

func main() {
	arguments := os.Args
	position := 1
	flag := false
	compteur := 0
	for range arguments {
		compteur++
	}

	if compteur >= 2 && arguments[1] == "--upper" {
		position = 2
		flag = true
	}

	for index, as := range arguments {
		if index >= position {
			num := BasicAtoi(as)
			if num == -1 {
				z01.PrintRune(' ')
			} else {
				if !flag {
					z01.PrintRune(rune('a' + num - 1))
				} else {
					z01.PrintRune(rune('A' + num - 1))
				}
			}
		}
	}
	z01.PrintRune(10)
}
