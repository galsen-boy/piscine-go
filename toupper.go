package piscine

func ToUpper(s string) string {
	bayefall := []rune(s)

	for i := 0; i < len(bayefall); i++ {
		if bayefall[i] >= 'a' && bayefall[i] <= 'z' {
			bayefall[i] -= 32
		}
	}

	return string(bayefall)
}
