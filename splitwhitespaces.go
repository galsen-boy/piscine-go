package piscine

func SplitWhiteSpaces(s string) []string {
	var daiba string = ""
	var answer []string
	for i := 0; i < len(s); i++ {
		if s[i] > ' ' {
			daiba = daiba + string(s[i])
		} else if daiba != "" {
			answer = append(answer, daiba)
			daiba = ""
		}
		if i == len(s)-1 {
			answer = append(answer, daiba)
			daiba = ""
		}

	}
	return answer
}
