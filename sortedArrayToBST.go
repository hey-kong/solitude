package main

// Leetcode 108. (easy)
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	i := len(nums) / 2
	return &TreeNode{
		Val:   nums[i],
		Left:  sortedArrayToBST(nums[:i]),
		Right: sortedArrayToBST(nums[i+1:]),
	}
}
