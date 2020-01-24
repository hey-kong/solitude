package main

// Leetcode 69. (easy)
func mySqrt(x int) int {
    if x == 0 {
		return 0
	}
	left := 1
	right := x
	for left <= right {
		mid := left + (right - left) / 2
		if mid <= x / mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left - 1
}
