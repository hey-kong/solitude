package main

// Leetcode 55. (medium)
func canJump(nums []int) bool {
	memo := make([]bool, len(nums))
	memo[len(nums)-1] = true
	for i := len(nums)-2; i >= 0; i-- {
		furthest := min(i+nums[i], len(nums)-1)
		for j := i+1; j <= furthest; j++ {
			if memo[j] {
				memo[i] = true
				break
			}
		}
	}
	return memo[0]
}

func canJump2(nums []int) bool {
	lastpos := len(nums) - 1
	for i := len(nums)-2; i >= 0; i-- {
		if i + nums[i] >= lastpos {
			lastpos = i
		}
	}
	return lastpos == 0
}
