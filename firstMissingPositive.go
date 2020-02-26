package main

// Leetcode 41. (hard)
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		for 1 <= nums[i] && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	i := 0
	for i < n {
		if nums[i] != i+1 {
			break
		}
		i++
	}
	return i + 1
}
