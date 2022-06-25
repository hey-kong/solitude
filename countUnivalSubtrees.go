package main

// Leetcode 250. (medium)
func countUnivalSubtrees(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res, _ := dfsCountUnivalSubtrees(root)
	return res
}

func dfsCountUnivalSubtrees(root *TreeNode) (int, bool) {
	if root.Left == nil && root.Right == nil {
		return 1, true
	}

	if root.Right == nil {
		lres, luni := dfsCountUnivalSubtrees(root.Left)
		if luni && root.Val == root.Left.Val {
			return lres + 1, true
		} else {
			return lres, false
		}

	}

	if root.Left == nil {
		rres, runi := dfsCountUnivalSubtrees(root.Right)
		if runi && root.Val == root.Right.Val {
			return rres + 1, true
		} else {
			return rres, false
		}
	}

	lres, luni := dfsCountUnivalSubtrees(root.Left)
	rres, runi := dfsCountUnivalSubtrees(root.Right)
	if luni && runi && root.Val == root.Left.Val && root.Val == root.Right.Val {
		return lres + rres + 1, true
	}
	return lres + rres, false
}
