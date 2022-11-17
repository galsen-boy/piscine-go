package piscine

func Any(f func(string) bool, daiba []string) bool {
	for _, v := range daiba {
		if f(v) {
			return true
		}
	}
	return false
}
