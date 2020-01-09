package main

// Leetcode 53. (easy)
func maxSubArray(nums []int) int {
	sum := nums[0]
	res := sum
	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		res = max(res, sum)
	}
	return res
}
