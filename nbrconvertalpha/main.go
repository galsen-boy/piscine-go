package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	l := 0
	for i := range args {
		l++
		i = i
	}
	cond := false
	if args[1] == "--upper" {
		cond = true
	}
	for j := 1; j < l; j++ {
		if args[j] == "1" {
			if cond == true {
				z01.PrintRune('A')
			} else {
				z01.PrintRune('a')
			}
		} else if args[j] == "2" {
			if cond == true {
				z01.PrintRune('B')
			} else {
				z01.PrintRune('b')
			}
		} else if args[j] == "3" {
			if cond == true {
				z01.PrintRune('C')
			} else {
				z01.PrintRune('c')
			}
		} else if args[j] == "4" {
			if cond == true {
				z01.PrintRune('D')
			} else {
				z01.PrintRune('d')
			}
		} else if args[j] == "5" {
			if cond == true {
				z01.PrintRune('E')
			} else {
				z01.PrintRune('e')
			}
		} else if args[j] == "6" {
			if cond == true {
				z01.PrintRune('F')
			} else {
				z01.PrintRune('f')
			}
		} else if args[j] == "7" {
			if cond == true {
				z01.PrintRune('G')
			} else {
				z01.PrintRune('g')
			}
		} else if args[j] == "8" {
			if cond == true {
				z01.PrintRune('H')
			} else {
				z01.PrintRune('h')
			}
		} else if args[j] == "9" {
			if cond == true {
				z01.PrintRune('I')
			} else {
				z01.PrintRune('i')
			}
		} else if args[j] == "10" {
			if cond == true {
				z01.PrintRune('J')
			} else {
				z01.PrintRune('j')
			}
		} else if args[j] == "11" {
			if cond == true {
				z01.PrintRune('K')
			} else {
				z01.PrintRune('k')
			}
		} else if args[j] == "12" {
			if cond == true {
				z01.PrintRune('L')
			} else {
				z01.PrintRune('l')
			}
		} else if args[j] == "13" {
			if cond == true {
				z01.PrintRune('M')
			} else {
				z01.PrintRune('m')
			}
		} else if args[j] == "14" {
			if cond == true {
				z01.PrintRune('N')
			} else {
				z01.PrintRune('n')
			}
		} else if args[j] == "15" {
			if cond == true {
				z01.PrintRune('O')
			} else {
				z01.PrintRune('o')
			}
		} else if args[j] == "16" {
			if cond == true {
				z01.PrintRune('P')
			} else {
				z01.PrintRune('p')
			}
		} else if args[j] == "17" {
			if cond == true {
				z01.PrintRune('Q')
			} else {
				z01.PrintRune('q')
			}
		} else if args[j] == "18" {
			if cond == true {
				z01.PrintRune('R')
			} else {
				z01.PrintRune('r')
			}
		} else if args[j] == "19" {
			if cond == true {
				z01.PrintRune('S')
			} else {
				z01.PrintRune('s')
			}
		} else if args[j] == "20" {
			if cond == true {
				z01.PrintRune('T')
			} else {
				z01.PrintRune('t')
			}
		} else if args[j] == "21" {
			if cond == true {
				z01.PrintRune('U')
			} else {
				z01.PrintRune('u')
			}
		} else if args[j] == "22" {
			if cond == true {
				z01.PrintRune('V')
			} else {
				z01.PrintRune('v')
			}
		} else if args[j] == "23" {
			if cond == true {
				z01.PrintRune('W')
			} else {
				z01.PrintRune('w')
			}
		} else if args[j] == "24" {
			if cond == true {
				z01.PrintRune('X')
			} else {
				z01.PrintRune('x')
			}
		} else if args[j] == "25" {
			if cond == true {
				z01.PrintRune('Y')
			} else {
				z01.PrintRune('y')
			}
		} else if args[j] == "26" {
			if cond == true {
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
