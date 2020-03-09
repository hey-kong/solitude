package main

// Leetcode 59. (medium)
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	num := 1
	up, down, left, right := 0, n-1, 0, n-1
	for {
		for i := left; i <= right; i++ {
			matrix[up][i] = num
			num++
		}
		up++
		if up > down {
			break
		}
		for i := up; i <= down; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		if right < left {
			break
		}
		for i := right; i >= left; i-- {
			matrix[down][i] = num
			num++
		}
		down--
		if down < up {
			break
		}
		for i := down; i >= up; i-- {
			matrix[i][left] = num
			num++
		}
		left++
		if left > right {
			break
		}
	}
	return matrix
}
