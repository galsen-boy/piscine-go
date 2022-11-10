package piscine

import (
	"github.com/01-edu/z01"
)

func IsSafe(queen int, fall int, pos [8]int) bool {
	for i := 0; i < queen; i++ {
		t := pos[i]
		if t == fall || t == fall-(queen-i) || t == fall+(queen-i) {
			return false
		}
	}
	return true
}

func solve(nb int, pos [8]int) {
	if nb == 8 {
		for i := 0; i < 8; i++ {
			z01.PrintRune(rune(pos[i] + 49))
		}
		z01.PrintRune(10)
	} else {
		for i := 0; i < 8; i++ {
			if IsSafe(nb, i, pos) {
				pos[nb] = i
				solve(nb+1, pos)
			}
		}
	}
}

func EightQueens() {
	pos := [8]int{0, 0, 0, 0, 0, 0, 0}
	solve(0, pos)
}
