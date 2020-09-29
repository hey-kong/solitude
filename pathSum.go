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

// Leetcode 113. (medium)
func pathSum2(root *TreeNode, sum int) [][]int {
	return recursivePathSum2(root, sum, []int{}, [][]int{})
}

func recursivePathSum2(root *TreeNode, sum int, cur []int, res [][]int) [][]int {
	if root == nil {
		return res
	}
	cur = append(cur, root.Val)
	if root.Val == sum && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		res = append(res, tmp)
	}
	res = recursivePathSum2(root.Left, sum-root.Val, cur, res)
	res = recursivePathSum2(root.Right, sum-root.Val, cur, res)
	cur = cur[:len(cur)-1]
	return res
}
