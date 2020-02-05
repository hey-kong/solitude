package main

// Leetcode 581. (easy)
func findUnsortedSubarray(nums []int) int {
	left, right := len(nums)-1, 0
	stack := []int{}
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			left = min(left, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = stack[:0]
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			right = max(right, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	if left >= right {
		return 0
	}
	return right - left + 1
}
