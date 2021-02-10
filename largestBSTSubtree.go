package main

// Leetcode 333. (medium)
func largestBSTSubtree(root *TreeNode) int {
	_, _, _, res := dfsLargestBSTSubtree(root)
	return res
}

func dfsLargestBSTSubtree(root *TreeNode) (bool, int, int, int) {
	if root == nil {
		return true, 1 << 31, -1 << 31, 0
	}
	if root.Left == nil && root.Right == nil {
		return true, root.Val, root.Val, 1
	}

	left, leftMin, leftMax, leftRes := dfsLargestBSTSubtree(root.Left)
	right, rightMin, rightMax, rightRes := dfsLargestBSTSubtree(root.Right)
	if !left || !right || root.Val <= leftMax || root.Val >= rightMin {
		return false, 0, 0, max(leftRes, rightRes)
	}
	if root.Left == nil {
		return true, root.Val, rightMax, rightRes + 1
	}
	if root.Right == nil {
		return true, leftMin, root.Val, leftRes + 1
	}
	return true, leftMin, rightMax, leftRes + rightRes + 1
}
