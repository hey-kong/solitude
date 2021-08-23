package main

// Leetcode 673. (medium)
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	dp, cnt := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		dp[i], cnt[i] = 1, 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					cnt[i] = cnt[j]
				} else if dp[j]+1 == dp[i] {
					cnt[i] += cnt[j]
				}
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := 0
	l := maxInt(dp...)
	for i := range dp {
		if dp[i] == l {
			res += cnt[i]
		}
	}
	return res
}
