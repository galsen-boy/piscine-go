package main

import (
	"os"
)

const (
	MaxInt64 = 9223372036854775807
	MinInt64 = -9223372036854775808
)

func ConvNbr(n int) string {
	donor := "0123456789"
	str := ""
	min := false
	if n < 0 {
		min = true
		n = n * -1
	}
	for {
		d := n % 10
		if d < 0 {
			d *= -1
		}
		str += string(donor[d])
		n /= 10
		if n == 0 {
			if min {
				str += "-"
			}
			break
		}
	}
	return StrRev1(str)
}

func StrRev1(s string) string {
	revStr := ""
	for i := len(s) - 1; i >= 0; i-- {
		revStr += string(rune(s[i]))
	}
	return revStr
}

func isSigne(str string, arr []string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}

func Atoi(s string) (int, bool) {
	if s == "" {
		return 0, false
	}
	n := 0
	s0 := s
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if len(s) < 1 {
			return 0, false
		}
	}
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 {
			return 0, false
		}
		if isOverfl(n*10, int(ch)) {
			return 0, false
		}
		n = n*10 + int(ch)
	}
	if s0[0] == '-' {
		n = -n
	}
	return n, true
}

func isOverfl(a int, b int) bool {
	if a > 0 && b > 0 && a+b < 0 {
		return true
	}
	if a < 0 && b < 0 && a+b > 0 {
		return true
	}
	return false
}

func main() {
	signe := []string{"+", "-", "/", "%", "*"}
	args := os.Args[1:]
	if len(args) != 3 {
	} else {
		if isSigne(args[1], signe) {
			nb1, nb1suc := Atoi(args[0])
			nb2, nb2suc := Atoi(args[2])
			// if (!success && nb1 == 0) || (nb2 == 0 && !success) {
			if !nb1suc || !nb2suc {
				return
			}
			switch args[1] {
			case signe[0]:
				// if nb1 > 0 && MaxInt64-nb1 <= nb2 || nb1 < 0 && MinInt64-nb1 >= nb2 {
				if isOverfl(nb1, nb2) {
					// os.Stdout.Write([]byte("Overflow"))
					return
				}
				os.Stdout.Write([]byte(ConvNbr(nb1+nb2) + "\n"))
			case signe[1]:
				if isOverfl(nb1, nb2) {
					// if MinInt64-nb1 < nb2 {
					// os.Stdout.Write([]byte("Overflow"))
					return
				}
				os.Stdout.Write([]byte(ConvNbr(nb1-nb2) + "\n"))
			case signe[2]:
				if nb2 == 0 {
					os.Stdout.Write([]byte("No division by 0" + "\n"))
				} else {
					os.Stdout.Write([]byte(ConvNbr(nb1/nb2) + "\n"))
				}
			case signe[3]:
				if nb2 == 0 {
					os.Stdout.Write([]byte("No modulo by 0" + "\n"))
				} else {
					os.Stdout.Write([]byte(ConvNbr(nb1%nb2) + "\n"))
				}
			case signe[4]:
				if nb1 == 0 || nb2 == 0 {
					os.Stdout.Write([]byte("0" + "\n"))
				} else {
					ans := nb1 * nb2
					if nb1 == ans/nb2 {
						os.Stdout.Write([]byte(ConvNbr(ans) + "\n"))
					} else {
						// os.Stdout.Write([]byte("Overflow"))
						return
					}
				}
			}
		} else {
			// os.Stdout.Write([]byte("\n"))
			return
		}
	}
}
