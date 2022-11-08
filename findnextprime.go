package piscine

func FindNextPrime(nb int) int {
	if nb < 3 {
		return 2
	}
	for baye := nb; ; baye++ {
		isprime := true
		for fall := 2; fall <= nb/2; fall++ {
			if baye%fall == 0 {
				isprime = false
				break
			}
		}
		if isprime {
			return baye
		}
	}
}
