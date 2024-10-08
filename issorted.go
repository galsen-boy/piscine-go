package piscine

func croissant(f func(a, b int) int, tab []int) bool {
	for i := 0; i < len(tab)-1; i++ {
		if f(tab[i], tab[i+1]) > 0 {
			return false
		}
	}
	return true
}

func decroissant(f func(a, b int) int, tab []int) bool {
	for i := len(tab) - 1; i > 0; i-- {
		if f(tab[i], tab[i-1]) > 0 {
			return false
		}
	}
	return true
}

func IsSorted(f func(a, b int) int, tab []int) bool {
	if !croissant(f, tab) && !decroissant(f, tab) {
		return false
	} else {
		return true
	}
}
