package piscine

func SortIntegerTable(table []int) {
	for j := 0; j < len(table)-1-i; j++ {
		if table[j] > table[j+1] {
			table[j+1], table[j] = table[j], table[j+1]
		}
	}
}
