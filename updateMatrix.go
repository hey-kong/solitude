package main

// Leetcode 542. (medium)
func updateMatrix(matrix [][]int) [][]int {
	mr, mc := len(matrix), len(matrix[0])
	for i := 0; i < mr; i++ {
		for j := 0; j < mc; j++ {
			if matrix[i][j] == 0 {
				continue
			}
			if i-1 >= 0 && j-1 >= 0 {
				matrix[i][j] = min(matrix[i-1][j], matrix[i][j-1]) + 1
			} else if i-1 >= 0 {
				matrix[i][j] = matrix[i-1][j] + 1
			} else if j-1 >= 0 {
				matrix[i][j] = matrix[i][j-1] + 1
			} else {
				matrix[i][j] = 10000
			}
		}
	}
	for i := mr - 1; i >= 0; i-- {
		for j := mc - 1; j >= 0; j-- {
			if matrix[i][j] == 0 {
				continue
			}
			if i+1 < mr && j+1 < mc {
				matrix[i][j] = min(matrix[i][j], min(matrix[i+1][j], matrix[i][j+1])+1)
			} else if i+1 < mr {
				matrix[i][j] = min(matrix[i][j], matrix[i+1][j]+1)
			} else if j+1 < mc {
				matrix[i][j] = min(matrix[i][j], matrix[i][j+1]+1)
			}
		}
	}
	return matrix
}
