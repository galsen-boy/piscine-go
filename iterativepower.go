package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	baye_fall := 1
	for i := 0; i < power; i++ {
		baye_fall = baye_fall * nb
	}
	return baye_fall
}
