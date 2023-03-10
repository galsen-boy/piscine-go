package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	BTreeApplyPostorder(root.Left, f)
	BTreeApplyPostorder(root.Right, f)
	f(root.Data)

	// if node = null
	//     return
	// preorder(node.left)
	// preorder(node.right)
	// visit(node)
}
