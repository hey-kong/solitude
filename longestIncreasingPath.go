package main

// Leetcode 329. (hard)
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	if len(matrix[0]) == 0 {
		return 0
	}

	res := 0
	memo := make([][]int, len(matrix))
	for i := range memo {
		memo[i] = make([]int, len(matrix[0]))
	}
	mr, mc := len(matrix), len(matrix[0])
	for i := 0; i < mr; i++ {
		for j := 0; j < mc; j++ {
			res = max(res, dfsLongestIncreasingPath(matrix, memo, i, j))
		}
	}
	return res
}

func dfsLongestIncreasingPath(matrix, memo [][]int, i, j int) int {
	mr, mc := len(matrix), len(matrix[0])
	if memo[i][j] != 0 {
		return memo[i][j]
	}
	up, right, down, left := 0, 0, 0, 0
	if i-1 >= 0 && matrix[i][j] < matrix[i-1][j] {
		up = dfsLongestIncreasingPath(matrix, memo, i-1, j)
	}
	if j+1 < mc && matrix[i][j] < matrix[i][j+1] {
		right = dfsLongestIncreasingPath(matrix, memo, i, j+1)
	}
	if i+1 < mr && matrix[i][j] < matrix[i+1][j] {
		down = dfsLongestIncreasingPath(matrix, memo, i+1, j)
	}
	if j-1 >= 0 && matrix[i][j] < matrix[i][j-1] {
		left = dfsLongestIncreasingPath(matrix, memo, i, j-1)
	}
	memo[i][j] = max(max(left, right), max(up, down)) + 1
	return memo[i][j]
}
