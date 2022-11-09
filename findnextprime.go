package piscine

func FindNextPrime(nb int) int {
	count := 0
	t := 0
	if nb == 0 || nb == 1 || nb < 0 {
		return 2
	}
	for i := nb - 1; i > 1; i-- {
		t = nb % i
		if t == 0 {
			count++
		}
	}
	if count >= 1 {
		return FindNextPrime(nb + 1)
	} else {
		return nb
	}
}
