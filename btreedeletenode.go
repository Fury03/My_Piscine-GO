package student

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	// Case 1: no left child
	if node.Left == nil {
		return BTreeTransplant(root, node, node.Right)
	}

	// Case 2: no right child
	if node.Right == nil {
		return BTreeTransplant(root, node, node.Left)
	}

	// Case 3: two children
	successor := BTreeMin(node.Right) // find minimum in right subtree
	if successor.Parent != node {
		// replace successor with its right child
		root = BTreeTransplant(root, successor, successor.Right)
		successor.Right = node.Right
		if successor.Right != nil {
			successor.Right.Parent = successor
		}
	}
	root = BTreeTransplant(root, node, successor)
	successor.Left = node.Left
	if successor.Left != nil {
		successor.Left.Parent = successor
	}

	return root
}
