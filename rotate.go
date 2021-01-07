package main

// Leetcode 48. (medium)
func rotateImage(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}

// Leetcode 189. (medium)
func rotate(nums []int, k int) {
	n := len(nums)
	if k >= n {
		k %= n
	}
	if k == 0 {
		return
	}

	i, val := 0, nums[0]
	start := i
	for _ = range nums {
		i += k
		if i >= n {
			i -= n
		}
		val, nums[i] = nums[i], val

		if i == start {
			i++
			val = nums[i]
			start = i
		}
	}
}
