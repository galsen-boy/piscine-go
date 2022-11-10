package piscine

import "github.com/01-edu/z01"

var x = []rune{'1', '2', '3', '4', '5', '6', '7', '8'}

func EightQueens() {
	print := true

	for i := 0; i < 8; i++ {
		for j := i + 1; j < 8; j++ {
			if i == j {
				continue
			}
			if x[i] == x[j] {
				print = false
				break
			}
			var diff rune
			if diff = x[i] - x[j]; x[i] < x[j] {
				diff = x[j] - x[i]
			}
			if diff == rune(j-i) {
				print = false
				break
			}
		}
		if !print {
			break
		}
	}

	if print {
		for _, r := range x {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}

	i := 7
	for i > 0 && x[i-1] >= x[i] {
		i--
	}

	if i <= 0 {
		return
	}

	j := 7

	for x[j] <= x[i-1] {
		j--
	}

	t := x[i-1]
	x[i-1] = x[j]
	x[j] = t

	j = 7
	for i < j {
		t := x[i]
		x[i] = x[j]
		x[j] = t
		i++
		j--
	}
	EightQueens()
}
