package main

import (
	"os"

	"github.com/01-edu/z01"
)

func check(x rune) bool {
	if x == 'a' || x == 'A' || x == 'e' || x == 'E' || x == 'o' || x == 'O' || x == 'u' || x == 'U' || x == 'i' || x == 'I' {
		return true
	}
	return false
}

func main() {
	args := os.Args[1:]
	rep := []rune{}
	ans := ""
	len := 0
	first := true
	for _, arg := range args {
		for _, j := range arg {
			if check(j) {
				rep = append(rep, j)
				len++
			}
		}
		if first {
			ans = arg
			first = false
			continue
		}
		ans = ans + " " + arg
	}
	cur := 0
	for _, c := range ans {
		if check(c) {
			z01.PrintRune(rep[len-cur-1])
			cur++
		} else {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
