package main

// Leetcode 1483. (hard)
type TreeAncestor struct {
	dp [][]int
}

func Constructor(n int, parent []int) TreeAncestor {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 1)
		dp[i][0] = parent[i]
	}
	j := 1
	for {
		finish := true
		for i := range dp {
			if dp[i][j-1] == -1 {
				dp[i] = append(dp[i], -1)
			} else {
				dp[i] = append(dp[i], dp[dp[i][j-1]][j-1])
				finish = false
			}
		}
		if finish {
			break
		}
		j++
	}
	return TreeAncestor{dp}
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	res, pos := node, 0
	for k != 0 && res != -1 {
		if pos >= len(this.dp[res]) {
			return -1
		}
		if k&1 != 0 {
			res = this.dp[res][pos]
		}
		k >>= 1
		pos += 1
	}
	return res
}
