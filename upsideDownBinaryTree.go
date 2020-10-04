package main

// Leetcode 156. (medium)
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	var parent, parentRight *TreeNode
	for root != nil {
		rootLeft := root.Left
		root.Left, parentRight = parentRight, root.Right
		root.Right, parent = parent, root
		root = rootLeft
	}
	return parent
}
