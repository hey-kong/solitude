package main

// Leetcode 861. (medium)
func matrixScore(A [][]int) int {
	m, n := len(A), len(A[0])
	res := 1 << (n - 1) * m
	for i := 1; i < n; i++ {
		ones := 0
		for j := 0; j < m; j++ {
			if A[j][i] == A[j][0] {
				ones++
			}
		}
		if ones > m-ones {
			res += 1 << (n - 1 - i) * ones
		} else {
			res += 1 << (n - 1 - i) * (m - ones)
		}
	}
	return res
}
