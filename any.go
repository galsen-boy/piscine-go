package piscine

func Any(f func(string) bool, arr []string) bool {
	for _, v := range arr {
		if f(v) {
			return true
		}
	}
	return false
}
