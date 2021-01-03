package main

import "strconv"

// Leetcode 257. (easy)
func binaryTreePaths(root *TreeNode) []string {
	return dfsBinaryTreePaths(root, "", []string{})
}

func dfsBinaryTreePaths(root *TreeNode, cur string, res []string) []string {
	if root == nil {
		return res
	}

	cur += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		res = append(res, cur)
		return res
	}

	cur += "->"
	if root.Left != nil {
		res = dfsBinaryTreePaths(root.Left, cur, res)
	}
	if root.Right != nil {
		res = dfsBinaryTreePaths(root.Right, cur, res)
	}
	cur = cur[:len(cur)-3]
	return res
}
