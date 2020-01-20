package main

// Leetcode 74. (medium)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	col := make([]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		col[i] = matrix[i][0]
	}
	left, right := 0, len(matrix)-1
	for left <= right {
		mid := left + (right - left) / 2
		if col[mid] == target {
			return true
		} else if col[mid] < target {
			left = mid + 1
		} else if target < col[mid] {
			right = mid - 1
		}
	}

	if left == 0 {
		return false
	}
	arr := matrix[left-1]
	left, right = 0, len(arr)-1
	for left <= right {
		mid := left + (right - left) / 2
		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			left = mid + 1
		} else if target < arr[mid] {
			right = mid - 1
		}
	}
	return false
}

// Leetcode 240. (medium)
func searchMatrix2(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	col := make([]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		col[i] = matrix[i][0]
	}
	left, right := 0, len(col)-1
	for left <= right {
		mid := left + (right - left) / 2
		if col[mid] == target {
			return true
		} else if col[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left == 0 {
		return false
	}
	down := right
	for i := 0; i < len(matrix); i++ {
		col[i] = matrix[i][len(matrix[i])-1]
	}
	left, right = 0, len(col)-1
	for left <= right {
		mid := left + (right - left) / 2
		if col[mid] == target {
			return true
		} else if col[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left == len(matrix) {
		return false
	}
	up := left
	for up <= down {
		left, right = 0, len(matrix[up])-1
		for left <= right {
			mid := left + (right - left) / 2
			if matrix[up][mid] == target {
				return true
			} else if matrix[up][mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		up++
	}
	return false
}

func searchMatrix3(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	i, j := len(matrix)-1, 0
	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			i--
		} else {
			j++
		}
	}
	return false
}
