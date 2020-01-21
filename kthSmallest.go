package main

// Leetcode 230. (medium)
func kthSmallest(root *TreeNode, k int) int {
	_, res := recursiveKthSmallest(root, 0, k, -1)
    return res
}

func recursiveKthSmallest(root *TreeNode, i, k, res int) (int, int) {
    if root == nil || i >= k {
		return i, res
	}
	i, res = recursiveKthSmallest(root.Left, i, k, res)
	i++
	if i == k {
		return i, root.Val
	}
	i, res = recursiveKthSmallest(root.Right, i, k, res)
	return i, res
}
