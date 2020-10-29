package main

// Leetcode 129. (medium)
func sumNumbers(root *TreeNode) int {
	return recursiveSumNumbers(root, 0, 0)
}

func recursiveSumNumbers(root *TreeNode, cur int, res int) int {
	if root == nil {
		return res
	}
	cur = cur*10 + root.Val
	if root.Left == nil && root.Right == nil {
		res = res + cur
	} else {
		res = recursiveSumNumbers(root.Left, cur, res)
		res = recursiveSumNumbers(root.Right, cur, res)
	}
	cur = (cur - root.Val) / 10
	return res
}
