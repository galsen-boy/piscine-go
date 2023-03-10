package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root != nil {
		// check the left and right branches is correct
		if root.Right != nil && root.Data > root.Right.Data {
			return false
		}
		if root.Left != nil && root.Data < root.Left.Data {
			return false
		}
		// recursive call
		if !BTreeIsBinary(root.Left) || !BTreeIsBinary(root.Right) {
			return false
		}
	}
	return true
}
