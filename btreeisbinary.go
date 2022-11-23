package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil || root.Right == nil || root.Left == nil {
		return true
	}
	daiba := false
	if root.Left.Data < root.Data && root.Right.Data > root.Data {
		daiba = BTreeIsBinary(root.Right)
		if daiba == true {
			daiba = BTreeIsBinary(root.Left)
		}
	}
	return daiba
}

package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Data > root.Data {
		return false
	}
	if root.Right != nil && root.Right.Data < root.Data {
		return false
	}
	if !BTreeIsBinary(root.Left) || !BTreeIsBinary(root.Right) {
		return false
	}

	return true
}