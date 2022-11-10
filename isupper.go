package piscine

func IsUpper(s string) bool {
	for _, bayefall := range s {
		if bayefall < 'A' || bayefall > 'Z' {
			return false
		}
	}
	return true
}
