package piscine

func SortSArr(s *[]string) {
	for i := 0; i < len((*s))-1; i++ {
		for j := 0; j < len((*s))-1-i; j++ {
			if (*s)[j] > (*s)[j+1] {
				(*s)[j], (*s)[j+1] = (*s)[j+1], (*s)[j]
			}
		}
	}
}

func SortWordArr(a []string) {
	SortSArr(&a)
}
