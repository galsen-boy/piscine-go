package piscine

func BasicAtoi2(s string) int {
	num := 0
	for _, i := range s {
		if i < '0' || i > '9' {
			return 0
		}
		num = num*10 + int(i) - '0'
	}
	return num
}
