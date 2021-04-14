package main

// Leetcode 198. (easy)
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	cur := nums[0]
	pre := 0
	for i := 1; i < len(nums); i++ {
		tmp := cur
		cur = max(pre+nums[i], cur)
		pre = tmp
	}
	return cur
}

// Leetcode 213. (medium)
func rob2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(rob2(nums[:n-1]), rob2(nums[1:]))
}

// Leetcode 337. (medium)
func rob3(root *TreeNode) int {
	res := recursiveRob3(root)
	return max(res[0], res[1])
}

func recursiveRob3(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left := recursiveRob3(root.Left)
	right := recursiveRob3(root.Right)
	chosen := root.Val + left[1] + right[1]
	notChosen := max(left[0], left[1]) + max(right[0], right[1])
	return []int{chosen, notChosen}
}
