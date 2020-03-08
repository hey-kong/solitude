package main

// Leetcode 103. (medium)
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	res = recursiveZigzagLevelOrder(root, 0, res)
	return res
}

func recursiveZigzagLevelOrder(root *TreeNode, i int, res [][]int) [][]int {
	if root == nil {
		return res
	}

	if len(res) == i {
		res = append(res, []int{})
	}
	if i%2 == 0 {
		res[i] = append(res[i], root.Val)
	} else {
		res[i] = append([]int{root.Val}, res[i]...)
	}
	res = recursiveZigzagLevelOrder(root.Left, i+1, res)
	res = recursiveZigzagLevelOrder(root.Right, i+1, res)
	return res
}
