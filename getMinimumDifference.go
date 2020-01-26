package main

import "math"

// Leetcode 530. (easy)
func getMinimumDifference(root *TreeNode) int {
	res := math.MaxInt32
	prev := -1
	_, res = recursiveGetMinimumDifference(root, prev, res)
	return res
}

func recursiveGetMinimumDifference(root *TreeNode, prev, res int) (int, int) {
    if root == nil {
		return prev, res
	}
	prev, res = recursiveGetMinimumDifference(root.Left, prev, res)
	if prev >= 0 {
		res = min(res, root.Val - prev)
	}
	prev = root.Val
	prev, res = recursiveGetMinimumDifference(root.Right, prev, res)
	return prev, res
}
