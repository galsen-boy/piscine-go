package piscine

func FindNextPrime(nb int) int {
	a := nb
	b := 0
	if !(nb < 0) {
		for i := nb; i >= 1; i-- {
			a = nb % i
			if a == 0 {
				b++
			}
		}
		if b == 2 {
			return nb
		} else {
			return FindNextPrime(nb + 1)
		}
	}
	return b
}
