package piscine

func BTreeLevelCount(root *TreeNode) int {
	var level int
	if root != nil {
		levelLeft := BTreeLevelCount(root.Left)
		levelRight := BTreeLevelCount(root.Right)
		if levelLeft > levelRight {
			level = levelLeft
		} else {
			level = levelRight
		}
		level++
	}
	return level
}

// func diver(root *TreeNode, level int) int {
// 	if root.Left == nil && root.Right == nil {
// 		return level
// 	} else {
// 		level++
// 		levelLeft := level
// 		levelRight := level
// 		// check branches and refresh levels value
// 		if root.Left != nil {
// 			levelLeft = diver(root.Left, level)
// 		}
// 		if root.Right != nil {
// 			levelRight = diver(root.Right, level)
// 		}
// 		// return max value
// 		if levelLeft > levelRight {
// 			return levelLeft
// 		} else {
// 			return levelRight
// 		}
// 	}
// }
