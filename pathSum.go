package main

// Leetcode 437. (easy)
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return recursivePathSum(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func recursivePathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Val == sum {
		res++
	}
	res += recursivePathSum(root.Left, sum-root.Val)
	res += recursivePathSum(root.Right, sum-root.Val)
	return res
}
