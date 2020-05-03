package main

// Leetcode 5401. (medium)
func kLengthApart(nums []int, k int) bool {
	left, right := -1, -1
	for i, num := range nums {
		if num == 1 {
			if left == -1 {
				left = i
			} else if right == -1 {
				right = i
				break
			}
		}
	}
	if left == -1 || right == -1 {
		return true
	}
	for i := right + 1; i < len(nums); i++ {
		if nums[i] == 1 {
			left, right = right, i
		}
		if right-left <= k {
			return false
		}
	}
	return true
}
