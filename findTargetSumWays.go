package main

// Leetcode 494. (medium)
func findTargetSumWays(nums []int, S int) int {
	// sum(P) - sum(N) = S
	// sum(P) - sum(N) + sum(P) + sum(N) = S + sum(nums)
	// 2 * sum(P) = S + sum(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < S || (S+sum)%2 == 1 {
		return 0
	}

	total := (S + sum) / 2
	dp := make([]int, total+1)
	dp[0] = 1
	for _, num := range nums {
		for j := total; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[total]
}
