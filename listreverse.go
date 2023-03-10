package piscine

func ListReverse(l *List) {
	var prev *NodeL = nil
	var current *NodeL = l.Head
	var next *NodeL = nil
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}
