package main

// Leetcode 487. (medium)
func findMaxConsecutiveOnes(nums []int) (res int) {
	dp0, dp1 := 0, 0
	for _, num := range nums {
		if num == 1 {
			dp0, dp1 = dp0+1, dp1+1
		} else {
			dp0, dp1 = 0, dp0+1
		}
		res = max(res, max(dp0, dp1))
	}
	return res
}
