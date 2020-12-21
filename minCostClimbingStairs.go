package main

// Leetcode 746. (easy)
func minCostClimbingStairs(cost []int) int {
	pre, cur := 0, 0
	for i := 2; i <= len(cost); i++ {
		pre, cur = cur, min(cur+cost[i-1], pre+cost[i-2])
	}
	return cur
}
