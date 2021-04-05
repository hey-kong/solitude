package main

// Leetcode 958. (medium)
func isCompleteTree(root *TreeNode) bool {
	queue := make([]*TreeNode, 1)
	queue[0] = root
	isEnd := false
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		if isEnd && cur != nil {
			return false
		}
		if cur == nil {
			isEnd = true
			continue
		}
		queue = append(queue, cur.Left)
		queue = append(queue, cur.Right)
	}
	return true
}
