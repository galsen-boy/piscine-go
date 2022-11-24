package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	daiba := BTreeLevelCount(root)
	for i := 1; i <= daiba; i++ {
		BtreeCurrentLevel(root, i, f)
	}
}

func BtreeCurrentLevel(root *TreeNode, daiba int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if daiba == 1 {
		f(root.Data)
	} else if daiba > 1 {
		BtreeCurrentLevel(root.Left, daiba-1, f)
		BtreeCurrentLevel(root.Right, daiba-1, f)
	}
}
