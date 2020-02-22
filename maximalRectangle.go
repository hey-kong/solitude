package main

// Leetcode 85. (hard)
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	mArea := 0
	heights := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				heights[j] = 0
			} else {
				heights[j]++
			}
		}
		// 'largestRectangleArea' refers to Leetcode 84.
		mArea = max(mArea, largestRectangleArea(heights))
	}
	return mArea
}
