package main

// Leetcode 5180. (hard)
func constrainedSubsetSum(nums []int, k int) int {
	res := nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	stack := []int{0}
	for i := 1; i < len(nums); i++ {
		dp[i] = nums[i] + max(0, dp[stack[0]])
		res = max(res, dp[i])
		if i-stack[0] == k {
			stack = stack[1:]
		}
		for len(stack) > 0 && dp[i] > dp[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
