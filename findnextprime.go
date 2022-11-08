package piscine

func FindNextPrime(nb int) int {
	for i := nb; i < 999; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}
