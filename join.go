package piscine

func Join(strs []string, sep string) string {
	i := ""
	for a, b := range strs {
		if a > 0 {
			i += sep + b
			continue
		}
		i += b
	}
	return i
}
