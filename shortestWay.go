package main

// Leetcode 1055. (medium)
func shortestWay(source string, target string) int {
	m, n := len(source), len(target)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, 26)
	}
	for i := 0; i < 26; i++ {
		dp[m][i] = m
	}

	for i := m - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			if source[i] == byte('a'+j) {
				dp[i][j] = i
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}

	res := 1
	tmp := 0
	for i := 0; i < n; i++ {
		c := target[i]
		if dp[0][c-'a'] == m {
			return -1
		}

		tmp = dp[tmp][c-'a']
		if tmp == m {
			res++
			tmp = dp[0][c-'a']
		}
		tmp++
	}
	return res
}
