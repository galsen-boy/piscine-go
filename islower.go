package piscine

func IsLower(s string) bool {
	for _, bayefall := range s {
		if bayefall < 'a' || bayefall > 'z' {
			return false
		}
	}
	return true
}
