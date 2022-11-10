package piscine

import "github.com/01-edu/z01"

const N = 8

var position = [N]int{}

func isSafe(queen_number, row_position int) bool {
	for i := 0; i < queen_number; i++ {
		other_row_pos := position[i]

		if other_row_pos == row_position || other_row_pos == row_position-(queen_number-i) || other_row_pos == row_position+(queen_number-i) {
			return false
		}
	}
	return true
}

func resoudre(a int) {
	if a == N {
		for i := 0; i < N; i++ {
			z01.PrintRune(rune(position[i] + '1'))
		}
		z01.PrintRune('\n')
	} else {
		for i := 0; i < N; i++ {
			if isSafe(a, i) {
				position[a] = i
				resoudre(a + 1)
			}
		}
	}
}

func EightQueens() {
	resoudre(0)
}
