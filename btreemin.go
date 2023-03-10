package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root.Left != nil {
		return BTreeMin(root.Left)
	} else {
		return root
	}
}
