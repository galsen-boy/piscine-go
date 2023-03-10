package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		levels := BTreeLevelCount(root)
		for i := 0; i < levels; i++ {
			applyOneLevel(root, i, f)
		}
	}
}

func applyOneLevel(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if level == 0 {
		f(root.Data)
	} else {
		applyOneLevel(root.Left, level-1, f)
		applyOneLevel(root.Right, level-1, f)
	}
}

// Write a function, BTreeApplyByLevel, that applies the function given by f,
// to each node of the tree given by root.

// test system output
// BTreeApplyByLevel(
// 	04
// 	├── 07
// 	│   ├── 12
// 	│   │   └── 10
// 	│   └── 05
// 	└── 01
// 		└── 02
// 			└── 03
// 	)
// 	 prints "0401020307051210" instead of "0401070205120310"
// 	exit status 1

// wrong vertical recursive, but live system test looks like requires a horizontal level prints
// func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
// 	if root != nil {
// 		f(root.Data)
// 		if root.Left != nil {
// 			BTreeApplyByLevel(root.Left, f)
// 		}
// 		if root.Right != nil {
// 			BTreeApplyByLevel(root.Right, f)
// 		}
// 	}
// }
