package piscine

func BasicAtoi(s string) int {
	x := 0
	for _, n := range s {
		y := 0
		for i := '1'; i <= n; i++ {
			y++
		}
		x = x*10 + y
	}
	return x
}
