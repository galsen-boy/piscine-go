package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	daiba := BTreeSearchItem(root, node.Data).Parent
	if daiba == nil {
		return rplc
	} else if node.Data < root.Data {
		root.Left = rplc
	} else {
		root.Right = rplc
	}
	rplc.Parent = daiba
	return root
}
