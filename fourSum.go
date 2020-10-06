package main

import "sort"

// Leetcode 18. (medium)
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	n := len(nums)
	res := [][]int{}
	sort.Ints(nums)

	for i := 0; i < n; i++ {
		if (i > 0 && nums[i] == nums[i-1]) || nums[i]+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}

		for j := i + 1; j < n; j++ {
			if (j > i+1 && nums[j] == nums[j-1]) || nums[i]+nums[j]+nums[n-1]+nums[n-2] < target {
				continue
			}

			last := n - 1
			for k := j + 1; k < n; k++ {
				if (k > j+1 && nums[k] == nums[k-1]) || nums[i]+nums[j]+nums[k]+nums[n-1] < target {
					continue
				}
				for k < last && nums[i]+nums[j]+nums[k]+nums[last] > target {
					last--
				}
				if k == last {
					break
				}
				if nums[i]+nums[j]+nums[k]+nums[last] == target {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[last]})
				}
			}
		}
	}
	return res
}
