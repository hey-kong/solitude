package main

// Leetcode 1746. (medium)
func maxSumAfterOperation(nums []int) int {
	dp0, dp1 := nums[0], nums[0]*nums[0]
	res := max(dp0, dp1)
	for i := 1; i < len(nums); i++ {
		dp1 = max(dp1+nums[i], max(dp0+nums[i]*nums[i], nums[i]*nums[i]))
		if dp0 < 0 {
			dp0 = nums[i]
		} else {
			dp0 += nums[i]
		}
		res = max(res, max(dp0, dp1))
	}
	return res
}
