package main

// Leetcode 463. (easy)
func islandPerimeter(grid [][]int) int {
	mr, mc := len(grid), len(grid[0])
	perimeter := 0
	for i := 0; i < mr; i++ {
		for j := 0; j < mc; j++ {
			if grid[i][j] == 0 {
				continue
			}
			adjCnt := 0
			if i-1 >= 0 && grid[i-1][j] == 1 {
				adjCnt++
			}
			if j-1 >= 0 && grid[i][j-1] == 1 {
				adjCnt++
			}
			perimeter += 4 - adjCnt*2
		}
	}
	return perimeter
}
