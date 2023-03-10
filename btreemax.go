package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root.Right != nil {
		return BTreeMax(root.Right)
	} else {
		return root
	}
}
