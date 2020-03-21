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

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			idx := getCircleIdx(i, j, n)
			out := 0
			for k := 0; k < idx; k++ {
				out += getCircleNumCnt(k, n)
			}
			cur := getCircleNumCnt(idx, n)
			if i == idx {
				matrix[i][j] = out + 1 + (j - idx)
			} else if j >= i {
				matrix[i][j] = out + (n - 2*idx) + (i - idx)
			} else if j == idx {
				matrix[i][j] = out + cur + 1 - (i - idx)
			} else {
				matrix[i][j] = out + cur/4*3 + 1 - (j - idx)
			}
		}
	}
	return matrix
}

func getCircleIdx(i, j, n int) int {
	return min(min(i, j), n-1-max(i, j))
}

func getCircleNumCnt(idx, n int) int {
	if res := (n - 2*idx - 1) * 4; res > 0 {
		return res
	}
	return 1
}
