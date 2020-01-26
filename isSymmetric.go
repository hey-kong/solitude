package main

// Leetcode 101. (easy)
func isSymmetric(root *TreeNode) bool {
    return recursiveIsSymmetric(root, root)
}

func recursiveIsSymmetric(node1, node2 *TreeNode) bool {
    if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return node1.Val == node2.Val && 
		recursiveIsSymmetric(node1.Left, node2.Right) && 
		recursiveIsSymmetric(node1.Right, node2.Left)
}
