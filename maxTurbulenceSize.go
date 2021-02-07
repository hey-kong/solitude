package main

// Leetcode 978. (medium)
func maxTurbulenceSize(A []int) (res int) {
	if len(A) <= 1 {
		return len(A)
	}

	dp1, dp2 := 1, 1
	for i := 1; i < len(A); i++ {
		if i%2 == 0 {
			if A[i-1] < A[i] {
				dp1 = dp1 + 1
				dp2 = 1
			} else if A[i-1] > A[i] {
				dp1 = 1
				dp2 = dp2 + 1
			} else {
				dp1 = 1
				dp2 = 1
			}
		} else {
			if A[i-1] > A[i] {
				dp1 = dp1 + 1
				dp2 = 1
			} else if A[i-1] < A[i] {
				dp1 = 1
				dp2 = dp2 + 1
			} else {
				dp1 = 1
				dp2 = 1
			}
		}
		res = max(res, max(dp1, dp2))
	}
	return res
}
