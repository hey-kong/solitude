package main

// Leetcode 977. (easy)
func sortedSquares(A []int) []int {
	i, j := 0, len(A)-1
	res := make([]int, len(A))
	k := len(A) - 1
	for i <= j {
		if abs(A[i]) > abs(A[j]) {
			res[k] = A[i] * A[i]
			i++
		} else {
			res[k] = A[j] * A[j]
			j--
		}
		k--
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
