package piscine

import "github.com/01-edu/z01"

const n int = 8

func EightQueens() {
	var rows [n]int
	getQueens(rows, 0)
}

func getQueens(rows [n]int, a int) {
	for i := 0; i < n; i++ {
		if checkPlace(rows, i, a) {
			rows[a] = i
			if a == n-1 {
				for j := 0; j < n; j++ {
					z01.PrintRune(getRuneQ(rows[j]))
				}
				z01.PrintRune('\n')
			} else {
				getQueens(rows, a+1)
			}
		}
	}
}

func checkPlace(rows [n]int, a, b int) bool {
	if b == 0 {
		return true
	}
	for i := 0; i < b; i++ {
		if rows[i] == a || rows[i] == a+b-i || rows[i] == a-b+i {
			return false
		}
	}
	return true
}

func getRuneQ(num int) (runenum rune) {
	switch num {
	case 0:
		return '1'
	case 1:
		return '2'
	case 2:
		return '3'
	case 3:
		return '4'
	case 4:
		return '5'
	case 5:
		return '6'
	case 6:
		return '7'
	default:
		return '8'
	}
}
