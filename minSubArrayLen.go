package main

// Leetcode 209. (medium)
func minSubArrayLen(s int, nums []int) int {
	left, right := 0, 0
	sum := 0
	res := 1<<31 - 1
	for right < len(nums) {
		sum += nums[right]
		right++
		for sum >= s {
			res = min(res, right-left)
			sum -= nums[left]
			left++
		}
	}
	if res == 1<<31-1 {
		return 0
	}
	return res
}
