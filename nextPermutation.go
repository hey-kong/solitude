package main

// Leetcode 31. (medium)
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i := len(nums) - 2
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
		for j := end; j >= i+1; j-- {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
}
