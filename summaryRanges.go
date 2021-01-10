package main

import "strconv"

// Leetcode 228. (easy)
func summaryRanges(nums []int) (res []string) {
	n := len(nums)
	i := 0
	for i < n {
		left := i
		for i = i + 1; i < n && nums[i-1]+1 == nums[i]; i++ {
		}

		if left == i-1 {
			res = append(res, strconv.Itoa(nums[left]))
		} else {
			res = append(res, strconv.Itoa(nums[left])+"->"+strconv.Itoa(nums[i-1]))
		}
	}
	return
}
