package main

// Leetcode 11. (medium)
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	area := 0
	for l < r {
		if height[l] < height[r] {
			area = max(area, height[l] * (r-l))
			l++
		} else {
			area = max(area, height[r] * (r-l))
			r--
		}
	}
	return area
}
