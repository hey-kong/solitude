package main

import "sort"

// Leetcode 16. (medium)
func threeSumClosest(nums []int, target int) int {
	best := 1<<31 - 1
	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}
			update(sum)
			if sum < target {
				for j+1 < k && nums[j] == nums[j+1] {
					j++
				}
				j++
			} else {
				for j < k-1 && nums[k-1] == nums[k] {
					k--
				}
				k--
			}
		}
	}
	return best
}
