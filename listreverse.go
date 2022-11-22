package piscine

func ListReverse(link *List) {
	daibanext := link.Head
	var daibaprev *NodeL
	link.Tail = link.Head

	for daibanext != nil {
		next := daibanext.Next
		daibanext.Next = daibaprev
		daibaprev = daibanext
		daibanext = next
	}

	link.Head = daibaprev
}
