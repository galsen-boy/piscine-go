package piscine

func StrRev(s string) string {
	temp := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	return string(temp)
}
