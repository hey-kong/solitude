package main

// Leetcode 104. (easy)
func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
