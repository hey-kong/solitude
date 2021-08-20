package main

// Leetcode 1230. (medium)
func probabilityOfHeads(prob []float64, target int) float64 {
	dp := make([]float64, target+1)
	dp[0] = 1
	for i, p := range prob {
		for j := min(target, i+1); j >= 1; j-- {
			dp[j] = p*dp[j-1] + (1-p)*dp[j]
		}
		dp[0] = (1 - p) * dp[0]
	}
	return dp[target]
}
