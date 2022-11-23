package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil || root.Right == nil || root.Left == nil {
		return true
	}
	checker := false
	if root.Left.Data < root.Data && root.Right.Data > root.Data {
		checker = BTreeIsBinary(root.Right)
		if checker == true {
			checker = BTreeIsBinary(root.Left)
		}
	}
	return checker
}
