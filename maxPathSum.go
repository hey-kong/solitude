package main

import "math"

// Leetcode 124. (hard)
func maxPathSum(root *TreeNode) int {
	_, maxSum := recursiveMaxPathSum(root)
	return maxSum
}

func recursiveMaxPathSum(root *TreeNode) (int, int) {
	if root == nil {
		return 0, math.MinInt32
	}
	left, leftMax := recursiveMaxPathSum(root.Left)
	right, rightMax := recursiveMaxPathSum(root.Right)

	maxPathSumWithRoot := root.Val + max(left, 0) + max(right, 0)
	maxSum := max(leftMax, rightMax)
	maxSum = max(maxSum, maxPathSumWithRoot)

	return root.Val + max(max(left, right), 0), maxSum
}
