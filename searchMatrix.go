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
