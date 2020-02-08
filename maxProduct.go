package main

// Leetcode 152. (medium)
func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp_max := make([]int, len(nums))
	dp_min := make([]int, len(nums))
	dp_max[0], dp_min[0] = nums[0], nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		preMax, preMin := dp_max[i-1], dp_min[i-1]
		if nums[i] < 0 {
			preMax, preMin = preMin, preMax
		}
		dp_max[i] = max(nums[i], preMax*nums[i])
		dp_min[i] = min(nums[i], preMin*nums[i])
		res = max(res, dp_max[i])
	}
	return res
}
