package main

// Leetcode 545. (medium)
func boundaryOfBinaryTree(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{root.Val}
	if root.Left == nil && root.Right == nil {
		return res
	}

	left := dfsLeft(root.Left, []int{})
	bottom := dfsBottom(root, []int{})
	right := dfsRight(root.Right, []int{})
	res = append(res, left...)
	res = append(res, bottom...)
	for i := len(right) - 1; i >= 0; i-- {
		res = append(res, right[i])
	}
	return res
}

func dfsLeft(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}

	if root.Left != nil || root.Right != nil {
		arr = append(arr, root.Val)
	}
	if root.Left != nil {
		arr = dfsLeft(root.Left, arr)
	} else if root.Right != nil {
		arr = dfsLeft(root.Right, arr)
	}
	return arr
}

func dfsRight(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}

	if root.Left != nil || root.Right != nil {
		arr = append(arr, root.Val)
	}
	if root.Right != nil {
		arr = dfsRight(root.Right, arr)
	} else if root.Left != nil {
		arr = dfsRight(root.Left, arr)
	}
	return arr
}

func dfsBottom(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}

	if root.Left == nil && root.Right == nil {
		arr = append(arr, root.Val)
	}
	if root.Left != nil {
		arr = dfsBottom(root.Left, arr)
	}
	if root.Right != nil {
		arr = dfsBottom(root.Right, arr)
	}
	return arr
}
