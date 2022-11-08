package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	for j := nb; j <= nb*2; j++ {
		for i := 2; i <= j/2+1; i++ {
			if j%i == 0 {
				break
			}
			if i == j/2+1 {
				return j
			}
		}
	}
	return 0
}
