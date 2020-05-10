package main

// Leetcode 5405. (medium)
func countTriplets(arr []int) int {
	dp := make([]int, len(arr))
	dp[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		dp[i] = dp[i-1] ^ arr[i]
	}
	res := 0
	for i := 0; i < len(arr)-1; i++ {
		for k := i + 1; k < len(arr); k++ {
			if (i == 0 && dp[k] == 0) || (i > 0 && dp[i-1] == dp[k]) {
				res += k - i
			}
		}
	}
	return res
}
