package main

// Leetcode 73. (medium)
func setZeroes(matrix [][]int) {
	setRow, setCol := false, false
	r, c := len(matrix), len(matrix[0])
	for i := 0; i < r; i++ {
		if matrix[i][0] == 0 {
			setRow = true
			break
		}
	}
	for i := 0; i < c; i++ {
		if matrix[0][i] == 0 {
			setCol = true
			break
		}
	}

	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if setRow {
		for i := 0; i < r; i++ {
			matrix[i][0] = 0
		}
	}
	if setCol {
		for i := 0; i < c; i++ {
			matrix[0][i] = 0
		}
	}
}
