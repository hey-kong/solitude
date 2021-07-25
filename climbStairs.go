package main

// Leetcode 70. (easy)
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	res := make([]int, n+1)
	res[1], res[2] = 1, 2
	for i := 3; i <= n; i++ {
		res[i] = res[i-2] + res[i-1]
	}
	return res[n]
}
