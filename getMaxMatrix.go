package main

// Leetcode m17.24. (hard)
func getMaxMatrix(matrix [][]int) []int {
	res := make([]int, 4)
	maxSum := matrix[0][0]
	for i := 0; i < len(matrix); i++ {
		tmpSum := make([]int, len(matrix[i]))
		for j := i; j < len(matrix); j++ {
			sum := 0
			start := 0
			for k := 0; k < len(matrix[j]); k++ {
				tmpSum[k] += matrix[j][k]
				if sum < 0 {
					sum = 0
					start = k
				}
				sum += tmpSum[k]
				if sum > maxSum {
					res = []int{i, start, j, k}
					maxSum = sum
				}
			}
		}
	}
	return res
}
