package main

// Leetcode 845. (medium)
func longestMountain(A []int) int {
	left, right := make([]int, len(A)), make([]int, len(A))
	for i := 1; i < len(A); i++ {
		if A[i-1] < A[i] {
			left[i] = left[i-1] + 1
		}
	}
	for i := len(A) - 2; i >= 0; i-- {
		if A[i] > A[i+1] {
			right[i] = right[i+1] + 1
		}
	}
	res := 0
	for i := 1; i < len(A)-1; i++ {
		if left[i] == 0 || right[i] == 0 {
			continue
		}
		res = max(res, left[i]+right[i]+1)
	}
	return res
}
