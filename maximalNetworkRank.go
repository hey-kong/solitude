package main

// Leetcode 1615. (medium)
func maximalNetworkRank(n int, roads [][]int) (res int) {
	inDegree := make([]int, n)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for _, road := range roads {
		matrix[road[0]][road[1]] = 1
		matrix[road[1]][road[0]] = 1
		inDegree[road[0]]++
		inDegree[road[1]]++
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			res = max(res, inDegree[i]+inDegree[j]-matrix[i][j])
		}
	}
	return
}
