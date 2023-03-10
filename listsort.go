package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	current := l
	if current != nil && current.Next == nil {
		return nil
	} else {
		for current != nil {
			index := current.Next
			for index != nil {
				if current.Data > index.Data {
					current.Data, index.Data = index.Data, current.Data
				}
				index = index.Next
			}
			current = current.Next
		}
	}
	return l
}
