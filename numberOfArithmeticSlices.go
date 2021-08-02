package main

// Leetcode 413. (medium)
func numberOfArithmeticSlices(nums []int) (res int) {
	if len(nums) < 2 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
			res += dp[i]
		}
	}
	return res
}
