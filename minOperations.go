package main

import "sort"

// Leetcode 5895. (medium)
func minOperations(grid [][]int, x int) int {
	m, n := len(grid), len(grid[0])
	arr := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (grid[i][j]-grid[0][0])%x != 0 {
				return -1
			}
			arr[i*n+j] = grid[i][j]
		}
	}

	sort.Ints(arr)
	target := 0
	for _, num := range arr {
		target += abs(num-arr[len(arr)/2]) / x
	}
	return target
}
