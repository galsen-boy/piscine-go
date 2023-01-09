package piscine

func ActiveBits(n int) int {
	daiba := 0
	for ; n > 1; n = n / 2 {
		daiba += n % 2
	}
	daiba += n
	return int(daiba)
}
