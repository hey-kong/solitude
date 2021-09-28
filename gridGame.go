package main

// Leetcode 2017. (medium)
func gridGame(grid [][]int) (res int64) {
	m, n := 2, len(grid[0])
	preSum := make([][]int64, m)
	for i := range preSum {
		preSum[i] = make([]int64, n+1)
		for j := 1; j <= n; j++ {
			preSum[i][j] = preSum[i][j-1] + int64(grid[i][j-1])
		}
	}

	res = 1<<63 - 1
	for i := 1; i <= n; i++ {
		res = minInt64(res, maxInt64(preSum[0][n]-preSum[0][i], preSum[1][i-1]))
	}
	return res
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
