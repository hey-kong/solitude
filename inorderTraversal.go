package main

// Leetcode 94. (medium)
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}
