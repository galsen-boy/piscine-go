package piscine

func Atoi(s string) int {
	x := 0
	z := 1
	for i, n := range s {
		y := 0
		if n < '0' || n > '9' {
			if n == '-' || n == '+' {
				if i != 0 {
					return 0
				}
				if n == '-' {
					z = -1
				}
			} else {
				return 0
			}
		}
		for i := '1'; i <= n; i++ {
			y++
		}
		x = x*10 + y*z
	}
	return x
}
