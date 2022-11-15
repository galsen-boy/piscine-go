package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, daiba := range a {
		for _, fall := range []rune(daiba) {
			z01.PrintRune(fall)
		}
		z01.PrintRune('\n')
	}
}
