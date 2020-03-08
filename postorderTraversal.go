package main

// Leetcode 145. (hard)
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	cur := root
	var last *TreeNode
	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			tmp := stack[len(stack)-1]
			if tmp.Right != nil && tmp.Right != last {
				cur = tmp.Right
			} else {
				res = append(res, tmp.Val)
				last = tmp
				stack = stack[:len(stack)-1]
			}
		}
	}
	return res
}
