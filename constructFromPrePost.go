package main

// Leetcode 889. (medium)
func constructFromPrePost(pre []int, post []int) *TreeNode {
	postMap := make(map[int]int, len(post))
	for i, num := range post {
		postMap[num] = i
	}
	return recursiveConstructFromPrePost(pre, 0, len(pre)-1, postMap, 0)
}

func recursiveConstructFromPrePost(pre []int, preL, preR int, postMap map[int]int, postL int) *TreeNode {
	if preL > preR {
		return nil
	}

	node := &TreeNode{
		Val:   pre[preL],
		Left:  nil,
		Right: nil,
	}
	if preL == preR {
		return node
	}
	i := postMap[pre[preL+1]]
	leftCount := i - postL + 1
	node.Left = recursiveConstructFromPrePost(pre, preL+1, preL+leftCount, postMap, postL)
	node.Right = recursiveConstructFromPrePost(pre, preL+leftCount+1, preR, postMap, i+1)
	return node
}
