package piscine

func AtoiBase(s string, base string) int {
	counter := 0
	sCounter := 0
	for a, b := range base {
		if b == '-' || b == '+' {
			return 0
		}
		counter++
		for j, b2 := range base {
			if b == b2 && a != j {
				return 0
			}
		}
	}
	for _, b := range s {
		if b == '-' || b == '+' {
			return 0
		}
		sCounter++
	}
	if counter < 2 {
		return 0
	}
	runes := []rune(s)
	result := 0
	j := 0
	for i := sCounter - 1; i >= 0; i-- {
		s := -1
		for k, b := range base {
			if b == runes[i] {
				s = k
			}
		}
		if s == (-1) {
			return 0
		}
		if j == 0 {
			result += s
		} else {
			daiba := 1
			for l := 0; l < j; l++ {
				daiba *= counter
			}
			result += daiba * s
		}
		j++
	}
	return result
}
