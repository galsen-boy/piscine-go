package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	counter := 0
	for _, a1 := range base {
		daiba := 0
		if a1 == '+' || a1 == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		for _, a2 := range base {
			if a1 == a2 {
				daiba++
			}
			if daiba == 2 {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
		counter++
	}
	if counter < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	mod := ""
	va := []rune(base)
	if nbr < 0 {
		z01.PrintRune('-')
	}
	for ; nbr != 0; nbr /= counter {
		index := nbr % counter
		if index < 0 {
			index = -index
		}
		mod = string(va[index]) + mod
	}
	for _, r := range mod {
		z01.PrintRune(r)
	}
}
