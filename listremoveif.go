package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	temp := l.Head
	var prev *NodeL = nil
	for temp != nil && temp.Data == data_ref {
		l.Head = temp.Next
		temp = l.Head
	}
	for temp != nil {
		for temp != nil && temp.Data != data_ref {
			prev = temp
			temp = temp.Next
		}
		if temp == nil {
			return
		}
		prev.Next = temp.Next
		temp = prev.Next
	}
}
