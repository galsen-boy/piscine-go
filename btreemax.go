package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else if root.Right == nil {
		return root
	}
	return BTreeMax(root.Right)
}
