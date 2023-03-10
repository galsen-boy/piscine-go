package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	pairs := len(a) - 1 // pairs in array
	sortedPairs := 0    // completely sorted at the moment
	for sortedPairs < pairs {
		for i := 0; i < pairs; i++ {
			// in task the ascending sorting is required, so 1 is bad response
			// fmt.Println("wtf ", f(a[i], a[i+1]))
			if f(a[i], a[i+1]) == 1 {
				a[i], a[i+1] = a[i+1], a[i]
				sortedPairs = 0
			} else {
				sortedPairs++
			}
		}
	}
}
