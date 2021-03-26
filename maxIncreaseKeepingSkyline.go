package main

// Leetcode 807. (medium)
func maxIncreaseKeepingSkyline(grid [][]int) (res int) {
	maxRow := make([]int, len(grid))
	maxCol := make([]int, len(grid[0]))

	for i := range grid {
		for j := range grid[i] {
			maxRow[i] = max(maxRow[i], grid[i][j])
			maxCol[j] = max(maxCol[j], grid[i][j])
		}
	}

	for i := range grid {
		for j := range grid[i] {
			res += min(maxRow[i], maxCol[j]) - grid[i][j]
		}
	}
	return res
}
