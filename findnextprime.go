package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 || nb > 10000000000 {
		return 2
	} else {
		prime := true
		for j := 2; j*j <= nb; j++ {
			if nb%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			return nb
		} else {
			return FindNextPrime(nb + 1)
		}
	}
}
