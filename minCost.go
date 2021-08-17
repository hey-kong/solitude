package main

import "math"

// Leetcode 256. (medium)
func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}

	dp := make([][]int, len(costs))
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	dp[0][0], dp[0][1], dp[0][2] = costs[0][0], costs[0][1], costs[0][2]
	for i := 1; i < len(costs); i++ {
		dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + costs[i][0]
		dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + costs[i][1]
		dp[i][2] = min(dp[i-1][0], dp[i-1][1]) + costs[i][2]
	}
	return min(dp[len(dp)-1][0], min(dp[len(dp)-1][1], dp[len(dp)-1][2]))
}

// Leetcode 265. (hard)
func minCostII(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}

	m, n := len(costs), len(costs[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	copy(dp[0], costs[0])
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if j == k {
					continue
				}
				if dp[i][j] == 0 {
					dp[i][j] = dp[i-1][k] + costs[i][j]
					continue
				}
				dp[i][j] = min(dp[i][j], dp[i-1][k]+costs[i][j])
			}
		}
	}
	return getMin(dp[m-1])
}

func getMin(a []int) int {
	res := math.MaxInt32
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}
