package piscine

func Index(s string, toFind string) int {
	a := []rune(s)
	b := []rune(toFind)
	index := -1
	for i := 0; i < len(a); i++ {
		if len(b)+i <= len(a) {
			if s[i:len(b)+i] == toFind {
				index = i
				break
			}
		}
	}
	return index
}
