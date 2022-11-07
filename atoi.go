package piscine

func Atoi(s string) int {
	if StrLen(s) == 0 {
		return 0
	}

	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if StrLen(s) < 1 {
			return 0
		}
	}

	nm := 0

	for _, i := range s {
		if i < '0' || i > '9' {
			return 0
		}
		nm = nm*10 + int(i)
	}

	return nm
}
