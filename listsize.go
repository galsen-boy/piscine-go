package piscine

func ListSize(l *List) int {
	var temp *NodeL = l.Head
	var count int = 0
	for temp != nil {
		// Visit to next node
		temp = temp.Next
		// Increase the node counter
		count++
	}
	return count
}
