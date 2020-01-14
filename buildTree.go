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

// Leetcode 106. (medium)
func buildTree2(inorder []int, postorder []int) *TreeNode {
	inMap := make(map[int]int)
	for i, num := range inorder {
		inMap[num] = i
	}
	return recursiveBuildTree2(postorder, 0, len(postorder)-1, inMap, len(inorder)-1)
}

func recursiveBuildTree2(postorder []int, postL, postR int, inMap map[int]int, inR int) *TreeNode {
	if postL > postR {
		return nil
	}

	node := &TreeNode{
		Val:   postorder[postR],
		Left:  nil,
		Right: nil,
	}
	i := inMap[node.Val]
	rightCount := inR - i
	node.Left = recursiveBuildTree2(postorder, postL, postR-rightCount-1, inMap, i-1)
	node.Right = recursiveBuildTree2(postorder, postR-rightCount, postR-1, inMap, inR)
	return node
}
