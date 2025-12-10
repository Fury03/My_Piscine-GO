package student

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Print the root first
	f(root.Data)

	// go left
	BTreeApplyPreorder(root.Left, f)

	// go right
	BTreeApplyPreorder(root.Right, f)
}
