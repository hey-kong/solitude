package main

// Leetcode 474. (medium)
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		zeros, ones := countZerosAndOnes(str)
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones]+1)
			}
		}
	}
	return dp[m][n]
}

func countZerosAndOnes(str string) (int, int) {
	zeros, ones := 0, 0
	for _, r := range str {
		if r == '0' {
			zeros++
		}
		if r == '1' {
			ones++
		}
	}
	return zeros, ones
}
