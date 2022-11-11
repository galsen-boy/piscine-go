package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	l := 0
	for i := range arguments {
		l++
		i = i
	}
	isupp := false
	if arguments[1] == "--upper" {
		isupp = true
	}
	for j := 1; j < l; j++ {
		if arguments[j] == "1" {
			if isupp == true {
				z01.PrintRune('A')
			} else {
				z01.PrintRune('a')
			}
		} else if arguments[j] == "2" {
			if isupp == true {
				z01.PrintRune('B')
			} else {
				z01.PrintRune('b')
			}
		} else if arguments[j] == "3" {
			if isupp == true {
				z01.PrintRune('C')
			} else {
				z01.PrintRune('c')
			}
		} else if arguments[j] == "4" {
			if isupp == true {
				z01.PrintRune('D')
			} else {
				z01.PrintRune('d')
			}
		} else if arguments[j] == "5" {
			if isupp == true {
				z01.PrintRune('E')
			} else {
				z01.PrintRune('e')
			}
		} else if arguments[j] == "6" {
			if isupp == true {
				z01.PrintRune('F')
			} else {
				z01.PrintRune('f')
			}
		} else if arguments[j] == "7" {
			if isupp == true {
				z01.PrintRune('G')
			} else {
				z01.PrintRune('g')
			}
		} else if arguments[j] == "8" {
			if isupp == true {
				z01.PrintRune('H')
			} else {
				z01.PrintRune('h')
			}
		} else if arguments[j] == "9" {
			if isupp == true {
				z01.PrintRune('I')
			} else {
				z01.PrintRune('i')
			}
		} else if arguments[j] == "10" {
			if isupp == true {
				z01.PrintRune('J')
			} else {
				z01.PrintRune('j')
			}
		} else if arguments[j] == "11" {
			if isupp == true {
				z01.PrintRune('K')
			} else {
				z01.PrintRune('k')
			}
		} else if arguments[j] == "12" {
			if isupp == true {
				z01.PrintRune('L')
			} else {
				z01.PrintRune('l')
			}
		} else if arguments[j] == "13" {
			if isupp == true {
				z01.PrintRune('M')
			} else {
				z01.PrintRune('m')
			}
		} else if arguments[j] == "14" {
			if isupp == true {
				z01.PrintRune('N')
			} else {
				z01.PrintRune('n')
			}
		} else if arguments[j] == "15" {
			if isupp == true {
				z01.PrintRune('O')
			} else {
				z01.PrintRune('o')
			}
		} else if arguments[j] == "16" {
			if isupp == true {
				z01.PrintRune('P')
			} else {
				z01.PrintRune('p')
			}
		} else if arguments[j] == "17" {
			if isupp == true {
				z01.PrintRune('Q')
			} else {
				z01.PrintRune('q')
			}
		} else if arguments[j] == "18" {
			if isupp == true {
				z01.PrintRune('R')
			} else {
				z01.PrintRune('r')
			}
		} else if arguments[j] == "19" {
			if isupp == true {
				z01.PrintRune('S')
			} else {
				z01.PrintRune('s')
			}
		} else if arguments[j] == "20" {
			if isupp == true {
				z01.PrintRune('T')
			} else {
				z01.PrintRune('t')
			}
		} else if arguments[j] == "21" {
			if isupp == true {
				z01.PrintRune('U')
			} else {
				z01.PrintRune('u')
			}
		} else if arguments[j] == "22" {
			if isupp == true {
				z01.PrintRune('V')
			} else {
				z01.PrintRune('v')
			}
		} else if arguments[j] == "23" {
			if isupp == true {
				z01.PrintRune('W')
			} else {
				z01.PrintRune('w')
			}
		} else if arguments[j] == "24" {
			if isupp == true {
				z01.PrintRune('X')
			} else {
				z01.PrintRune('x')
			}
		} else if arguments[j] == "25" {
			if isupp == true {
				z01.PrintRune('Y')
			} else {
				z01.PrintRune('y')
			}
		} else if arguments[j] == "26" {
			if isupp == true {
				z01.PrintRune('Z')
			} else {
				z01.PrintRune('z')
			}
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(10)
}
