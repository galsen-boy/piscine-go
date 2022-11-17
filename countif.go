package piscine

func CountIf(f func(string) bool, tab []string) int {
	compteur := 0
	for _, v := range tab {
		if f(v) {
			compteur++
		}
	}
	return compteur
}
