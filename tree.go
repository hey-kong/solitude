package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNode2 struct {
	Val    int
	Left   *TreeNode2
	Right  *TreeNode2
	Parent *TreeNode2
}

type TreeNode3 struct {
	Val   int
	Left  *TreeNode3
	Right *TreeNode3
	Next  *TreeNode3
}
