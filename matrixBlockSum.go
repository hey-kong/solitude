package main

// Leetcode 1314. (medium)
func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	preSum := make([][]int, m+1)
	for i := range preSum {
		preSum[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1] + mat[i-1][j-1]
		}
	}

	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = getOfMatrixBlockSum(preSum, i+1+k, j+1+k, m, n) -
				getOfMatrixBlockSum(preSum, i-k, j+1+k, m, n) -
				getOfMatrixBlockSum(preSum, i+1+k, j-k, m, n) +
				getOfMatrixBlockSum(preSum, i-k, j-k, m, n)
		}
	}
	return res
}

func getOfMatrixBlockSum(preSum [][]int, x, y int, m, n int) int {
	i, j := min(max(x, 0), m), min(max(y, 0), n)
	return preSum[i][j]
}
