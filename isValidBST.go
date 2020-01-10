package main

import "math"

// Leetcode 98. (medium)
func isValidBST(root *TreeNode) bool {
	return recursiveIsValidBST(root, math.MinInt64, math.MaxInt64)
}

func recursiveIsValidBST(root *TreeNode, low, high int) bool {
	if root == nil {
		return true
	}

	if root.Val <= low || root.Val >= high {
		return false
	}

	return recursiveIsValidBST(root.Left, low, root.Val) && recursiveIsValidBST(root.Right, root.Val, high)
}
