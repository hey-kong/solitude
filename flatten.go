package main

// Leetcode 114. (medium)
func flatten(root *TreeNode)  {
	recursiveFlatten(root)
}

func recursiveFlatten(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = recursiveFlatten(root.Left)
	root.Right = recursiveFlatten(root.Right)

	left := root.Left
	if left == nil {
		return root
	}

	for left.Right != nil {
		left = left.Right
	}
	left.Right = root.Right
	root.Right = root.Left
	root.Left = nil
	return root
}
