package main

// Leetcode 918. (medium)
func maxSubarraySumCircular(A []int) int {
	if len(A) == 1 {
		return A[0]
	}

	sum := A[0]
	sum1, res1 := A[0], A[0]
	for i := 1; i < len(A); i++ {
		sum += A[i]
		if sum1 < 0 {
			sum1 = A[i]
		} else {
			sum1 += A[i]
		}
		res1 = max(res1, sum1)
	}
	sum2, res2 := A[1], A[1]
	for i := 2; i < len(A)-1; i++ {
		if sum2 > 0 {
			sum2 = A[i]
		} else {
			sum2 += A[i]
		}
		res2 = min(res2, sum2)
	}
	return max(res1, sum-res2)
}
