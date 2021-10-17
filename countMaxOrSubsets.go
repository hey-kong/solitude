package main

// Leetcode 5904. (medium)
func countMaxOrSubsets(nums []int) int {
	res, cnt := 0, 0
	n := (1 << len(nums)) - 1
	for i := 1; i <= n; i++ {
		cur := 0
		for j := 0; j < len(nums); j++ {
			if (i>>j)&1 != 0 {
				cur |= nums[j]
			}
		}
		if cur > res {
			res = cur
			cnt = 1
		} else if cur == res {
			cnt++
		}
	}
	return cnt
}
