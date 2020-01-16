package main

// Leetcode 45. (hard)
func jump(nums []int) int {
	memo := make([]int, len(nums))
	for i := 0; i < len(memo)-1; i++ {
		memo[i] = -1
	}
	for i := len(nums)-2; i >= 0; i-- {
		furthest := min(i+nums[i], len(nums)-1)
		for j := i+1; j <= furthest; j++ {
			if memo[j] == -1 {
				continue
			}
			if memo[i] == -1 {
				memo[i] = memo[j]+1
			} else {
				memo[i] = min(memo[j]+1, memo[i])
			}
		}
	}
}

func jump2(nums []int) int {
	steps := 0
	maxpos := 0
	end := 0
	for i := 0; i < len(nums)-1; i++ {
		maxpos = max(maxpos, nums[i]+i)
		if i == end {
			if end == maxpos {
				return -1
			}
			end = maxpos
			steps++
		}
	}
	return steps
}
