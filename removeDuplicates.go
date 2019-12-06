package main

// Leetcode 26. (easy)
func removeDuplicates(nums []int) int {
	i := 0
	for _, num := range nums {
		if nums[i] != num {
			i++
			nums[i] = num
		}
	}
	return i + 1
}
