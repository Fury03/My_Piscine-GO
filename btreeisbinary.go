package student

func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, "", "")
}

func isBST(node *TreeNode, min string, max string) bool {
	if node == nil {
		return true
	}

	// check if node value is valid with min/max limits
	if (min != "" && node.Data <= min) || (max != "" && node.Data >= max) {
		return false
	}

	// go left with updated max
	// go right with updated min
	return isBST(node.Left, min, node.Data) &&
		isBST(node.Right, node.Data, max)
}
