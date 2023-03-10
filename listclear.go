package piscine

func ListClear(l *List) {
	// Remove linked list node one by one
	for l.Head != nil {
		// Visit to next node
		l.Head = l.Head.Next
	}
}
