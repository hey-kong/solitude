package main

// Leetcode 1162. (medium)
func maxDistance(grid [][]int) int {
	N := len(grid)
	queue := [][2]int{}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	// no land
	if len(queue) == 0 {
		return -1
	}

	dx := [4]int{-1, 0, 1, 0}
	dy := [4]int{0, 1, 0, -1}
	res := -1
	for len(queue) != 0 {
		locate := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			x := locate[0] + dx[i]
			y := locate[1] + dy[i]
			if x >= 0 && x < N && y >= 0 && y < N && grid[x][y] == 0 {
				grid[x][y] = grid[locate[0]][locate[1]] + 1
				queue = append(queue, [2]int{x, y})
				res = max(res, grid[x][y]-1)
			}
		}
	}
	return res
}
