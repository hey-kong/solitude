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

// Leetcode 189. (easy)
func rotate(nums []int, k int) {
	if k >= len(nums) {
		k %= len(nums)
	}
	if k == 0 {
		return
	}

	i := 0
	start := 0
	tmp := nums[i]
	for _ = range nums {
		i += k
		if i >= len(nums) {
			i -= len(nums)
		}
		tmp, nums[i] = nums[i], tmp

		if i == start {
			i++
			start = i
			tmp = nums[i]
		}
	}
}
