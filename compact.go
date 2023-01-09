package piscine

func Compact(ptr *[]string) int {
	daiba := *ptr

	compteur := 0

	for _, v := range *ptr {
		if v != "" {
			daiba[compteur] = v
			compteur++
		}
	}
	*ptr = daiba[0:compteur]
	return compteur
}
