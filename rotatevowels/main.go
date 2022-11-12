package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Len(r interface{}) int {
	count := 0
	switch a := r.(type) {
	case []rune:
		for range a {
			count++
		}
	case []string:
		for range a {
			count++
		}
	case string:
		for range a {
			count++
		}
	}
	return count
}

func main() {
	var a1, a2, rev []rune
	for _, arg := range os.Args[1:] {
		for _, k := range arg {
			if containsRune("aeiouAEIOU", k) {
				a1 = append(a1, k)
			}
		}
	}
	for i := Len(a1) - 1; i >= 0; i-- {
		rev = append(rev, a1[i])
	}

	m := 0
	for i, arg := range os.Args[1:] {
		for _, j := range arg {
			if containsRune("aeiouAEIOU", j) {
				a2 = append(a2, rev[m])
				m++
			} else {
				a2 = append(a2, j)
			}
		}
		if i != Len(os.Args)-1 {
			a2 = append(a2, ' ')
		}
	}
	PrintStr(string(a2))
	z01.PrintRune('\n')
}

func PrintStr(str string) {
	for _, r := range str {
		z01.PrintRune(r)
	}
}

func containsRune(s string, r rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}
