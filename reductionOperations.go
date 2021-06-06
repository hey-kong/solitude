package main

import "sort"

// Leetcode 5777. (medium)
func reductionOperations(nums []int) (sum int) {
	sort.Ints(nums)
	res := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			res[i] = res[i-1]
		} else {
			res[i] = res[i-1] + 1
		}
		sum += res[i]
	}
	return sum
}
