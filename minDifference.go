package main

// Leetcode 5790. (medium)
func minDifference(nums []int, queries [][]int) []int {
	pre := make([][]int, len(nums)+1)
	for i := range pre {
		pre[i] = make([]int, 101)
	}
	for i := 1; i <= len(nums); i++ {
		copy(pre[i], pre[i-1])
		pre[i][nums[i-1]]++
	}

	res := make([]int, len(queries))
	for k, q := range queries {
		left, right := q[0], q[1]
		last, ans := 0, 100
		for i := 1; i <= 100; i++ {
			if pre[left][i] != pre[right+1][i] {
				if last != 0 {
					ans = min(ans, i-last)
				}
				last = i
			}
		}
		if ans != 100 {
			res[k] = ans
		} else {
			res[k] = -1
		}
	}
	return res
}
