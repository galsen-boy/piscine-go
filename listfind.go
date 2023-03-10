package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	n := l.Head
	for n != nil {
		if comp(n.Data, ref) {
			break
		}
		n = n.Next
	}
	return &n.Data
}
