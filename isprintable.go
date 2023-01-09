package piscine

func IsPrintable(s string) bool {
	for _, bayefall := range s {
		if bayefall < ' ' {
			return false
		}
	}
	return true
}
