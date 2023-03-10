package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	// create new node
	// n := &NodeL{Data: data, Next: l.Head}
	// Make next of new Node as next of prev_node
	// n.Next = l.Head
	// make next of prev_node as new_node
	// l.Head = n

	// code above is revert directed syntax

	// from task
	// 1 2 3
	n := &NodeL{Data: data}

	// 4
	if l.Head == nil {
		l.Head = &NodeL{Data: data}
		return
	}

	// 4 step too, looks like mistyping
	n.Next = nil

	// 5
	last := l.Head
	for last.Next != nil {
		last = last.Next
	}

	// 6
	last.Next = n

	// https://www.geeksforgeeks.org/linked-list-set-2-inserting-a-node/
	// C# "append" method , from global function included all methods
}
