package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	daiba := 0
	for _, v1 := range nbr {
		for i2, v2 := range baseFrom {
			if v1 == v2 {
				daiba = daiba*StringLen(baseFrom) + i2
			}
		}
	}

	x := ""
	for daiba != 0 {

		x = string(baseTo[daiba%StringLen(baseTo)]) + x
		daiba /= StringLen(baseTo)
	}

	return x
}

func StringLen(str string) int {
	i := 0
	for range str {
		i++
	}
	return i
}
