package piscine

import (
	"strconv"
)

func BasicAtoi2(s string) int {
	i, valeur := strconv.Atoi(s)

	if valeur == nil {
		return i
	} else {
		return 0
	}
}
