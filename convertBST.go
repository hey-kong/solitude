package main

// Leetcode 538. (easy)
func convertBST(root *TreeNode) *TreeNode {
	recursiveConvertBST(root, 0)
	return root
}

func recursiveConvertBST(root *TreeNode, num int) int {
	if root == nil {
		return num
	}

	num = recursiveConvertBST(root.Right, num)
	root.Val += num
	num = root.Val
	num = recursiveConvertBST(root.Left, num)
	return num
}
