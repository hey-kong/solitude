package main

// Leetcode 5791. (medium)
func countSubIslands(grid1 [][]int, grid2 [][]int) (res int) {
	if len(grid1) == 0 || len(grid1[0]) == 0 {
		return
	}

	for i := 0; i < len(grid1); i++ {
		for j := 0; j < len(grid1[0]); j++ {
			if grid2[i][j] == 1 && dfsCountSubIslands(grid1, grid2, i, j, true) {
				res++
			}
		}
	}
	return
}

func dfsCountSubIslands(grid1 [][]int, grid2 [][]int, i, j int, sub bool) bool {
	if i < 0 || i >= len(grid1) || j < 0 || j >= len(grid1[0]) || grid2[i][j] == 0 {
		return sub
	}

	if grid1[i][j] == 0 {
		sub = false
	}
	grid2[i][j] = 0
	sub = dfsCountSubIslands(grid1, grid2, i+1, j, sub)
	sub = dfsCountSubIslands(grid1, grid2, i, j+1, sub)
	sub = dfsCountSubIslands(grid1, grid2, i-1, j, sub)
	sub = dfsCountSubIslands(grid1, grid2, i, j-1, sub)
	return sub
}
