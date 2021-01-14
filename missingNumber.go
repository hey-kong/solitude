package main

// Leetcode 268. (easy)
func missingNumber(nums []int) int {
	res := len(nums)
	for i, num := range nums {
		res ^= i ^ num
	}
	return res
}

// Leetcode m53 - II. (easy)
func missingNumber2(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == mid {
			left = mid + 1
		} else if nums[mid] > mid {
			right = mid - 1
		}
	}
	return left
}
