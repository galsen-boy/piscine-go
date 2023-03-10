package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListPushFront(l *List, data interface{}) {
	// create new node
	n := &NodeL{Data: data, Next: l.Head}
	// Make next of new Node as next of prev_node
	n.Next = l.Head
	// make next of prev_node as new_node
	l.Head = n
}
