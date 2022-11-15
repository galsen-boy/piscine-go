package piscine

func MakeRange(min, max int) []int {
	if max-min <= 0 {
		var answer []int
		return answer
	}
	answer := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		answer[i] = i + min
	}
	return answer
}
