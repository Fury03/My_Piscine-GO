package student

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// queue to store nodes
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// remove first element
		node := queue[0]
		queue = queue[1:]

		// apply the function to the node data
		f(node.Data)

		// add left child if exists
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		// add right child if exists
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}
