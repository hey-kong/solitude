package main

// Leetcode 200. (medium)
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	mr, mc := len(grid), len(grid[0])
	num := 0
	for i := 0; i < mr; i++ {
		for j := 0; j < mc; j++ {
			if grid[i][j] == '1' {
				num++
				grid = dfsNumIslands(grid, i, j)
			}
		}
	}
	return num
}

func dfsNumIslands(grid [][]byte, r, c int) [][]byte {
	mr, mc := len(grid), len(grid[0])
	if r == -1 || r == mr || c == -1 || c == mc || grid[r][c] == '0' {
		return grid
	}
	grid[r][c] = '0'
	grid = dfsNumIslands(grid, r-1, c)
	grid = dfsNumIslands(grid, r, c+1)
	grid = dfsNumIslands(grid, r+1, c)
	grid = dfsNumIslands(grid, r, c-1)
	return grid
}
