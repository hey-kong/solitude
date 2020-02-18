package main

// Leetcode 75. (medium)
func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}
	left, right := 0, len(nums)-1
	i := 0
	for i <= right {
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
			i++
		} else if nums[i] == 1 {
			i++
		} else {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		}
	}
	return
}
