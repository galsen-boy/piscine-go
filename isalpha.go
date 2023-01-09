package piscine

func IsAlpha(s string) bool {
	for _, bayefall := range s {
		if bayefall < '0' || bayefall > '9' &&
			bayefall < 'A' || bayefall > 'Z' &&
			bayefall < 'a' || bayefall > 'z' {
			return false
		}
	}
	return true
}
