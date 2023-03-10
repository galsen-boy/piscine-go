package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// replacement := node
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left { // node is left child of parent
		node.Parent.Left = rplc // set left child of replacement parent to rplc
	} else {
		node.Parent.Right = rplc // set right child of replacement parent to rplc
	}
	if rplc != nil {
		rplc.Parent = node.Parent // using replacement parent node, connect replacement branch to node parent position
	}

	return root
}

// https://www.cs.dartmouth.edu/~thc/cs10/lectures/0428/0428.html
// search on page ... protected void transplant

// wrong in live system testing, but ok in example
// func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
// 	if root != nil {
// 		foundNode := BTreeSearchItem(root, node.Data)
// 		foundNode.Data = rplc.Data
// 	}
// 	return root
// }
