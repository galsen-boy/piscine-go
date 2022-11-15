package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		var answer []int
		return answer
	}
	answer := make([]int, min, max)
	for i := 0; i < max-min; i++ {
		answer[i] = i + min
	}
	return answer
}
