package main

// Leetcode 283. (easy)
func moveZeroes(nums []int)  {
	if len(nums) <= 1 {
		return
	}

	mark := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[mark] = nums[mark], nums[i]
			mark++
		}
	}
}
