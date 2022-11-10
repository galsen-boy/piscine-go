package piscine

func Capitalize(s string) string {
	a := []rune(s)
	b := 0
	for i := range a {
		if (a[i] >= 'a' && a[i] <= 'z') || (a[i] >= 'A' && a[i] <= 'Z') || (a[i] <= '9' && a[i] >= '0') {
			b++
		} else {
			b = 0
		}

		if b == 1 && a[i] >= 'a' && a[i] <= 'z' {
			a[i] -= 32
		} else if b > 1 && a[i] >= 'A' && a[i] <= 'Z' {
			a[i] += 32
		}
	}
	s = string(a)
	return s
}
