package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	a := '0'
	b := '1'
	c := '2'
	for i := a; i <= 55; i++ {
		for j := b; j <= 56; j++ {
			for k := c; k <= 57; k++ {
				if i < j && j < k {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					if i != 55 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
