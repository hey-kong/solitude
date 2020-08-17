package main

// Leetcode 110. (easy)
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(recursiveIsBalanced(root.Left)-recursiveIsBalanced(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func recursiveIsBalanced(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(recursiveIsBalanced(root.Left), recursiveIsBalanced(root.Right)) + 1
}
