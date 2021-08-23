package main

import "sort"

// Leetcode 368. (medium)
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dp := make([]int, len(nums))
	dp[0] = 1
	maxLen, maxVal := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if maxLen < dp[i] {
			maxLen, maxVal = dp[i], nums[i]
		}
	}

	res := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		if dp[i] == maxLen && maxVal%nums[i] == 0 {
			res = append(res, nums[i])
			maxLen--
			maxVal = nums[i]
		}
	}
	return res
}
