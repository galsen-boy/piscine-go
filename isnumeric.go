package piscine

func IsNumeric(s string) bool {
	for _, bayefall := range s {
		if bayefall < '0' || bayefall > '9' {
			return false
		}
	}
	return true
}
