package main

// Leetcode 978. (medium)
func maxTurbulenceSize(arr []int) int {
	res := 1
	n := len(arr)
	dp1, dp2 := 1, 1
	for i := 1; i < n; i++ {
		if arr[i-1] > arr[i] {
			dp1, dp2 = dp2+1, 1
		} else if arr[i-1] < arr[i] {
			dp1, dp2 = 1, dp1+1
		} else {
			dp1, dp2 = 1, 1
		}
		res = max(res, max(dp1, dp2))
	}
	return res
}
