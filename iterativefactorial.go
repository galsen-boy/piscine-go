package piscine

func IterativeFactorial(nb int) int {
	if nb < 20 {
		var a int
		a = 1
		if nb < 0 {
			return 0
		}
		for l := 1; l <= nb; l++ {
			a = a * l
		}
		return a
	} else {
		return 0
	}
}

func main() {
}
