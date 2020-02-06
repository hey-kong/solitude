package main

// Leetcode 31. (medium)
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i := len(nums) - 2
	if nums[i] < nums[i+1] {
		nums[i], nums[i+1] = nums[i+1], nums[i]
		return
	}
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i < 0 {
		bubbleSort(nums, 0, len(nums)-1)
		return
	}
	j := len(nums) - 1
	for nums[i] >= nums[j] {
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	bubbleSort(nums, i+1, len(nums)-1)
}

func bubbleSort(nums []int, start, end int) {
	for i := start; i <= end-1; i++ {
		for j := i + 1; j <= end; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
