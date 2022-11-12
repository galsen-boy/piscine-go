package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args

	if len(arguments) <= 1 {
		z01.PrintRune('\n')
		return
	}

	namePro := arguments[0]
	count := 0
	str := ""
	var vowels []rune
	var ind []int
	for range arguments {
		count++
	}
	for i := 0; i < count-1; i++ {
		if arguments[i] != namePro {
			str = str + arguments[i] + " "
		}
	}
	str = str + arguments[count-1]
	for i, a := range str {
		if isVowel(a) {
			vowels = append(vowels, a)
			ind = append(ind, i)

		}
	}
	n := 0
	runes := []rune(str + " ")
	for range runes {
		n++
	}
	m := 0
	for range vowels {
		m++
	}
	for i := 0; i < n; i++ {
		for j := range ind {
			if i == ind[j] {
				runes[i] = vowels[m-j-1]
			}
		}
	}
	for i := range runes {
		z01.PrintRune(runes[i])
	}
	z01.PrintRune('\n')
}

func isVowel(a rune) bool {
	return (a == 'a' || a == 'e' || a == 'u' || a == 'o' || a == 'i' || a == 'A' || a == 'E' || a == 'U' || a == 'O' || a == 'I')
}
