package main

import "sort"

// Leetcode 646. (medium)
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})

	dp := make([]int, len(pairs))
	dp[0] = 1
	for i := 1; i < len(pairs); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if pairs[j][1] < pairs[i][0] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return maxInt(dp...)
}
