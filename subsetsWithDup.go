package main

import "sort"

// Leetcode 90. (medium)
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	res = recursiveSubsetsWithDup(0, len(nums)-1, nums, []int{}, res)
	return res
}

func recursiveSubsetsWithDup(start, end int, nums, tmp []int, res [][]int) [][]int {
	res = append(res, tmp)
	for i := start; i <= end; i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		cur := make([]int, len(tmp))
		copy(cur, tmp)
		cur = append(cur, nums[i])
		res = recursiveSubsetsWithDup(i+1, end, nums, cur, res)
	}
	return res
}
