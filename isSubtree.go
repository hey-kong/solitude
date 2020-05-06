package main

// Leetcode 572. (easy)
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	return isSametree(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSametree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	return s.Val == t.Val && isSametree(s.Left, t.Left) && isSametree(s.Right, t.Right)
}
