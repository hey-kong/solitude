package main

// Leetcode 983. (medium)
func mincostTickets(days []int, costs []int) int {
	last := days[len(days)-1]
	dp := make([]int, last+1)
	idx := 0
	for i := 1; i < last+1; i++ {
		if i != days[idx] {
			dp[i] = dp[i-1]
			continue
		}
		cost := 1<<31 - 1
		oneDaysAgo := i - 1
		sevenDaysAgo := i - 7
		thirtyDaysAgo := i - 30
		if sevenDaysAgo < 0 {
			sevenDaysAgo = 0
		}
		if thirtyDaysAgo < 0 {
			thirtyDaysAgo = 0
		}

		cost = min(cost, dp[oneDaysAgo]+costs[0])
		cost = min(cost, dp[sevenDaysAgo]+costs[1])
		cost = min(cost, dp[thirtyDaysAgo]+costs[2])
		dp[i] = cost
		idx++
	}
	return dp[last]
}
