package main

import "sort"

// Leetcode m16.24. (medium)
func pairSums(nums []int, target int) (res [][]int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] == target {
			res = append(res, []int{nums[i], nums[j]})
			i++
			j--
		} else if nums[i]+nums[j] > target {
			j--
		} else if nums[i]+nums[j] < target {
			i++
		}
	}
	return
}
