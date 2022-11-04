package piscine

import "github.com/01-edu/z01"

func PrintNum(k int) {
	i := '0'
	if k == 0 {
		z01.PrintRune('0')
	}
	for j := 1; j <= k%10; j++ {
		i++
	}
	for j := -1; j >= k%10; j-- {
		i++
	}
	if k/10 != 0 {
		PrintNbr(k / 10)
	}
	z01.PrintRune(i)
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	PrintNum(n)
}
