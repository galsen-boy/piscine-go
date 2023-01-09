package piscine

func IterativeFactorial(nb int) int {
	baye_fall := 1
	if nb > 25 || nb < 0 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		baye_fall = baye_fall * i
	}
	return baye_fall
}
