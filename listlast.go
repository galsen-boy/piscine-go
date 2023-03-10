package piscine

func ListLast(l *List) interface{} {
	// check the nodes
	if l.Head == nil {
		return nil
	}
	n := l.Head
	for n != nil {
		if n.Next == nil {
			return n.Data
		}
		n = n.Next
	}
	return nil
}
