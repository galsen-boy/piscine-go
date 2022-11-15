package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		var answer []int
		return answer
	}
	answer := make([]int, min, max)
	for i := min; i < max; i++ {
		answer[i] = i + min
	}
	return answer
}
