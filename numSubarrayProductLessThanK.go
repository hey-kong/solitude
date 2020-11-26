package main

// Leetcode 713. (medium)
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	i, j := 0, 0
	res := 0
	p := 1
	for j < len(nums) {
		p *= nums[j]
		for p >= k {
			res += j - i
			p /= nums[i]
			i++
		}
		j++
	}
	res += (1 + (j - i)) * (j - i) / 2
	return res
}
