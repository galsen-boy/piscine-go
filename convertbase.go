package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	ResBase := 0
	for _, v1 := range nbr {
		for i2, v2 := range baseFrom {
			if v1 == v2 {
				ResBase = ResBase*StringLen(baseFrom) + i2
			}
		}
	}

	x := ""
	for ResBase != 0 {

		x = string(baseTo[ResBase%StringLen(baseTo)]) + x
		ResBase /= StringLen(baseTo)
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
