package main

var result [][]int

// Leetcode 102. (medium)
func levelOrder(root *TreeNode) [][]int {
	result = make([][]int, 0)
	recursiveLevelOrder(root, 0)
	return result
}

func recursiveLevelOrder(root *TreeNode, level int) {
	if root == nil {
		return
	}

	if level == len(result) {
		arr := make([]int, 0)
		result = append(result, arr)
	}

	result[level] = append(result[level], root.Val)

	recursiveLevelOrder(root.Left, level+1)
	recursiveLevelOrder(root.Right, level+1)
}
