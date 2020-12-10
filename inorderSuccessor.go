package main

// Leetcode 510. (medium)
func inorderSuccessor(node *TreeNode2) *TreeNode2 {
	if node.Right != nil {
		node = node.Right
		for node.Left != nil {
			node = node.Left
		}
		return node
	}

	for node.Parent != nil && node.Parent.Right == node {
		node = node.Parent
	}
	return node.Parent
}
