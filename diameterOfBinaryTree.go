package main

// Leetcode 543. (easy)
func diameterOfBinaryTree(root *TreeNode) int {
	_, maxPath := recursiveDiameterOfBinaryTree(root)
	return maxPath
}

func recursiveDiameterOfBinaryTree(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	left, leftMax := recursiveDiameterOfBinaryTree(root.Left)
	right, rightMax := recursiveDiameterOfBinaryTree(root.Right)
	maxPath := max(leftMax, rightMax)
	maxPath = max(maxPath, left + right)

	return max(left, right) + 1, maxPath
}
