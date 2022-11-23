package piscine

func BTreeLevelCount(root *TreeNode) int {
	compteur := 0
	if root == nil {
		return compteur
	}
	compteur++
	counterright := compteur + BTreeLevelCount(root.Right)
	counterleft := compteur + BTreeLevelCount(root.Left)
	if counterright > counterleft {
		return counterright
	} else {
		return counterleft
	}
}
