package piscine

func Map(f func(int) bool, a []int) []bool {
	daiba := []bool{}
	for _, v := range a {
		daiba = append(daiba, f(v))
	}
	return daiba
}
