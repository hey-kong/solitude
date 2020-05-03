package main

// Leetcode 456. (medium)
func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	minArr := make([]int, len(nums))
	minArr[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		minArr[i] = min(minArr[i-1], nums[i])
	}
	stack := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > minArr[i] {
			for len(stack) != 0 && minArr[i] >= stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) != 0 && nums[i] > stack[len(stack)-1] {
				return true
			}
			stack = append(stack, nums[i])
		}
	}
	return false
}
