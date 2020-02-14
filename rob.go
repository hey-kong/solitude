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
