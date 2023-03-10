package piscine

func ListMerge(l1 *List, l2 *List) {
	n := l2.Head
	for n != nil {
		ListPushBack(l1, n.Data)
		n = n.Next
	}
}
