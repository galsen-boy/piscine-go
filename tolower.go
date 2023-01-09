package piscine

func ToLower(s string) string {
	bayefall := []rune(s)

	for i := 0; i < len(bayefall); i++ {
		if bayefall[i] >= 'A' && bayefall[i] <= 'Z' {
			bayefall[i] += 32
		}
	}

	return string(bayefall)
}
