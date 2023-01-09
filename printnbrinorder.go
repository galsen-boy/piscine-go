package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var digits []int
	if n == 0 {
		z01.PrintRune('0')
	}
	counter := 0
	for n > 0 {
		digits = append(digits, n%10+48)
		n = n / 10
		counter++
	}
	for i := 0; i < counter-1; i++ {
		for j := i + 1; j < counter; j++ {
			if digits[i] > digits[j] {
				digits[i], digits[j] = digits[j], digits[i]
			}
		}
	}
	for i := 0; i < counter; i++ {
		z01.PrintRune(rune(digits[i]))
	}
}
