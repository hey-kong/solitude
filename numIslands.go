package main

// Leetcode 200. (medium)
func numIslands(grid [][]byte) (res int) {
	if len(grid) == 0 {
		return
	}

	mr, mc := len(grid), len(grid[0])
	for i := 0; i < mr; i++ {
		for j := 0; j < mc; j++ {
			if grid[i][j] == '1' {
				res++
				dfsNumIslands(grid, i, j)
			}
		}
	}
	return
}

func dfsNumIslands(grid [][]byte, r, c int) {
	mr, mc := len(grid), len(grid[0])
	if r == -1 || r == mr || c == -1 || c == mc || grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	dfsNumIslands(grid, r-1, c)
	dfsNumIslands(grid, r, c+1)
	dfsNumIslands(grid, r+1, c)
	dfsNumIslands(grid, r, c-1)
}

func numIslands2(grid [][]byte) (res int) {
	if len(grid) == 0 {
		return 0
	}

	dx := []int{-1, 0, 1, 0}
	dy := []int{0, -1, 0, 1}
	mr, mc := len(grid), len(grid[0])
	queue := [][2]int{}
	for i := 0; i < mr; i++ {
		for j := 0; j < mc; j++ {
			if grid[i][j] == '1' {
				res++
				grid[i][j] = '0'
				queue = append(queue, [2]int{i, j})
				for len(queue) > 0 {
					x, y := queue[0][0], queue[0][1]
					queue = queue[1:]
					for k := 0; k < 4; k++ {
						tmpX, tmpY := x+dx[k], y+dy[k]
						if tmpX >= 0 && tmpX < mr && tmpY >= 0 && tmpY < mc && grid[tmpX][tmpY] == '1' {
							grid[tmpX][tmpY] = '0'
							queue = append(queue, [2]int{tmpX, tmpY})
						}
					}
				}
			}
		}
	}
	return
}
