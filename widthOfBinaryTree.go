package main

// Leetcode 662. (medium)
var widthOfBinaryTreeRes int
var widthOfBinaryTreeMap map[int]int

func widthOfBinaryTree(root *TreeNode) int {
	widthOfBinaryTreeRes, widthOfBinaryTreeMap = 0, make(map[int]int)
	recursiveWidthOfBinaryTree(root, 0, 0)
	return widthOfBinaryTreeRes
}

func recursiveWidthOfBinaryTree(root *TreeNode, depth, pos int) {
	if root == nil {
		return
	}
	if _, ok := widthOfBinaryTreeMap[depth]; !ok {
		widthOfBinaryTreeMap[depth] = pos
	}
	widthOfBinaryTreeRes = max(widthOfBinaryTreeRes, pos-widthOfBinaryTreeMap[depth]+1)
	recursiveWidthOfBinaryTree(root.Left, depth+1, 2*pos)
	recursiveWidthOfBinaryTree(root.Right, depth+1, 2*pos+1)
}
