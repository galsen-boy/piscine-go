package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	if n > len(s) || n < 1 {
		return 0
	}
	return runes[n-1]
}
