package main

import (
	"os"

	"github.com/01-edu/z01"
)

func BasicAtoi(s string) int {
	d := 0

	for _, i := range s {
		if '0' <= i && i <= '9' {
			pam := 0
			for x := '1'; x <= i; x++ {
				pam = pam + 1
			}

			d = d*10 + pam

		} else {
			d = -1
			break
		}
	}

	if (d != -1) && !(1 <= d && d <= 26) {
		d = -1
	}

	return d
}

func main() {
	arguments := os.Args
	position := 1
	flag := false
	compteur := 0
	for range arguments {
		compteur++
	}

	if compteur >= 2 && arguments[1] == "--upper" {
		position = 2
		flag = true
	}

	for index, as := range arguments {
		if index >= position {
			num := BasicAtoi(as)
			if num == 1 {
				z01.PrintRune(' ')
			} else {
				if !flag {
					z01.PrintRune(rune('a' + num - 1))
				} else {
					z01.PrintRune(rune('A' + num - 1))
				}
			}
		}
	}
	z01.PrintRune(10)
}
