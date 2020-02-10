package main

// Leetcode 695. (medium)
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	maxArea := 0
	mr, mc := len(grid), len(grid[0])
	for i := 0; i < mr; i++ {
		for j := 0; j < mc; j++ {
			if grid[i][j] == 1 {
				area := 0
				grid, area = dfsMaxAreaOfIsland(grid, i, j, area)
				maxArea = max(maxArea, area)
			}
		}
	}
	return maxArea
}

func dfsMaxAreaOfIsland(grid [][]int, r, c, area int) ([][]int, int) {
	mr, mc := len(grid), len(grid[0])
	if r == -1 || r == mr || c == -1 || c == mc || grid[r][c] == 0 {
		return grid, area
	}
	grid[r][c] = 0
	area++
	grid, area = dfsMaxAreaOfIsland(grid, r-1, c, area)
	grid, area = dfsMaxAreaOfIsland(grid, r, c+1, area)
	grid, area = dfsMaxAreaOfIsland(grid, r+1, c, area)
	grid, area = dfsMaxAreaOfIsland(grid, r, c-1, area)
	return grid, area
}
