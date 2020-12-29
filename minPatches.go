package main

// Leetcode 330. (hard)
func minPatches(nums []int, n int) int {
	i := 0
	sum := 1
	res := 0
	for sum <= n {
		if i >= len(nums) || sum < nums[i] {
			res++
			sum = 2 * sum
		} else {
			sum += nums[i]
			i++
		}
	}
	return res
}
