package main

// Leetcode 5699. (medium)
func countRestrictedPaths(n int, edges [][]int) int {
	g1 := make([][]int, n+1)
	g2 := make([][]bool, n+1)
	for i := 1; i <= n; i++ {
		g1[i] = make([]int, n+1)
		g2[i] = make([]bool, n+1)
		for j := 1; j <= n; j++ {
			g1[i][j] = 0x3f3f3f3f
			if i == j {
				g1[i][j] = 0
			}
		}
	}
	for _, e := range edges {
		g1[e[0]][e[1]] = e[2]
		g1[e[1]][e[0]] = e[2]
	}

	inDegree := make([]int, n+1)
	v := make([]bool, n+1)
	for j := 0; j < n; j++ {
		tmp := 0
		for i := 1; i <= n; i++ {
			if !v[i] && (tmp == 0 || g1[n][i] < g1[n][tmp]) {
				tmp = i
			}
		}

		v[tmp] = true

		for i := 1; i <= n; i++ {
			g1[n][i] = min(g1[n][i], g1[n][tmp]+g1[tmp][i])
		}
	}

	for _, e := range edges {
		if g1[n][e[0]] > g1[n][e[1]] {
			g2[e[0]][e[1]] = true
			inDegree[e[1]]++
		} else if g1[n][e[0]] < g1[n][e[1]] {
			g2[e[1]][e[0]] = true
			inDegree[e[0]]++
		}
	}
	dp := make([]int, n+1)
	dp[1] = 1
	queue := []int{}
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		tmp := queue[0]
		queue = queue[1:]
		for i := 1; i <= n; i++ {
			if g2[tmp][i] {
				inDegree[i]--
				if inDegree[i] == 0 {
					queue = append(queue, i)
				}
				dp[i] = (dp[i] + dp[tmp]) % 1000000007
			}
		}
	}
	return dp[n]
}
