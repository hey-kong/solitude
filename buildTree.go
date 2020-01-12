package main

// Leetcode 105. (medium)
func buildTree(preorder []int, inorder []int) *TreeNode {
	inMap := make(map[int]int, len(inorder))
	for i, num := range inorder {
		inMap[num] = i
	}
	return recursiveBuildTree(preorder, 0, len(preorder)-1, inMap, 0)
}

func recursiveBuildTree(preorder []int, preL, preR int, inMap map[int]int, inL int) *TreeNode {
	if preL > preR {
		return nil
	}

	i := inMap[preorder[preL]]
	leftCount := i - inL
	node := &TreeNode{
		Val:   preorder[preL],
		Left:  nil,
		Right: nil,
	}
	node.Left = recursiveBuildTree(preorder, preL+1, preL+leftCount, inMap, inL)
	node.Right = recursiveBuildTree(preorder, preL+leftCount+1, preR, inMap, i+1)
	return node
}
