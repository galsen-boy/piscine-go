package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	// fordelete := BTreeSearchItem(root, node.Data)
	root = deleteRec(root, node.Data)

	return root
}

/* A recursive function todelete an existing "data" in BST*/
func deleteRec(root *TreeNode, data string) *TreeNode {
	/* Base Case: If the tree is empty */
	if root == nil {
		return root
	}

	/* Otherwise, recur down the tree */
	if data < root.Data {
		root.Left = deleteRec(root.Left, data)
	} else if data > root.Data {
		root.Right = deleteRec(root.Right, data)
		// if key is same as root's
		// key, then This is the
		// node to be deleted
	} else {
		// node with only one child or no child
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// node with two children: Get the inorder
		// successor (smallest in the right subtree)
		root.Data = minValue(root.Right)

		// Delete the inorder successor
		root.Right = deleteRec(root.Right, root.Data)
	}

	return root
}

func minValue(root *TreeNode) string {
	minv := root.Data
	for root.Left != nil {
		minv = root.Left.Data
		root = root.Left
	}
	return minv
}

// https://www.geeksforgeeks.org/binary-search-tree-set-2-delete/
// javascript section
