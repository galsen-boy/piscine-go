package piscine

func TrimAtoi(s string) int {
	bayefall := ""
	for _, touba := range s {
		if touba >= '0' && touba <= '9' || touba == '-' {
			if len(bayefall) == 0 && touba == '-' {
				bayefall += string(touba)
				continue
			}
			if touba != '-' {
				bayefall += string(touba)
			}
		}
	}
	return Atoi(bayefall)
}
