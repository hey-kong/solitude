package main

var res []int

// Leetcode 199. (medium)
func rightSideView(root *TreeNode) []int {
	res = []int{}
	recursiveRightSideView(root, 0)
	return res
}

func recursiveRightSideView(root *TreeNode, deep int) {
	if root == nil {
		return
	}

	if len(res) == deep {
		res = append(res, root.Val)
	} else {
		res[deep] = root.Val
	}
	recursiveRightSideView(root.Left, deep+1)
	recursiveRightSideView(root.Right, deep+1)
}
