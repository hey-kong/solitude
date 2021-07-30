package main

// Leetcode 1014. (medium)
func maxScoreSightseeingPair(values []int) int {
	res := 0
	pre := values[0] + 0
	for j := 1; j < len(values); j++ {
		res = max(res, values[j]-j+pre)
		pre = max(pre, values[j]+j)
	}
	return res
}
