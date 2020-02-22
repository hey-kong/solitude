package main

// Leetcode 84. (hard)
func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	stack := make([]int, 1)
	stack[0] = 0
	mArea := 0
	for i := 1; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - (stack[len(stack)-1] + 1)
			}
			mArea = max(mArea, height*width)
		}
		stack = append(stack, i)
	}
	return mArea
}
