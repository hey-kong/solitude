package main

// Leetcode 89. (medium)
func grayCode(n int) []int {
	res := []int{0}
	for i := 1; i <= n; i++ {
		res = append(res, res...)
		inc := len(res) / 2
		left, right := inc-1, inc
		for left >= 0 {
			res[right] = res[left] + inc
			left--
			right++
		}
	}
	return res
}
