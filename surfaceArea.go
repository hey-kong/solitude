package main

// Leetcode 892. (easy)
func surfaceArea(grid [][]int) int {
	direction := make([][2]int, 4)
	direction[0], direction[1], direction[2], direction[3] = [2]int{-1, 0}, [2]int{0, -1}, [2]int{1, 0}, [2]int{0, 1}

	N := len(grid)
	res := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] > 0 {
				res += 2
			}
			for _, d := range direction {
				if i+d[0] < 0 || i+d[0] >= N || j+d[1] < 0 || j+d[1] >= N {
					res += grid[i][j]
					continue
				}
				res += max(grid[i][j]-grid[i+d[0]][j+d[1]], 0)
			}
		}
	}
	return res
}
