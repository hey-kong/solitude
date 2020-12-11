package main

// Leetcode 376. (medium)
func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	up := make([]int, len(nums))
	down := make([]int, len(nums))
	up[0], down[0] = 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			up[i] = max(up[i-1], down[i-1]+1)
			down[i] = down[i-1]
		} else if nums[i-1] > nums[i] {
			up[i] = up[i-1]
			down[i] = max(down[i-1], up[i-1]+1)
		} else {
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
	}
	return max(up[len(nums)-1], down[len(nums)-1])
}
