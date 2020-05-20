package main

import "sort"

// Leetcode 611. (medium)
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			target := nums[i] + nums[j]
			left, right := j+1, len(nums)-1
			for left <= right {
				mid := left + (right-left)/2
				if nums[mid] >= target {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
			res += right - j
		}
	}
	return res
}
