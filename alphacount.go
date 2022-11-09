package piscine

func AlphaCount(s string) int {
	counter := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 || s[i] >= 97 && s[i] <= 122 {
			counter++
		}
	}
	return counter
}
