package main

// Leetcode 5753. (hard)
func largestPathValue(colors string, edges [][]int) int {
	in := make([]int, len(colors))
	g := make([][]int, len(colors))
	for i := range edges {
		in[edges[i][1]]++
		g[edges[i][0]] = append(g[edges[i][0]], edges[i][1])
	}

	q := []int{}
	for i := range in {
		if in[i] == 0 {
			q = append(q, i)
		}
	}

	cnt := 0
	res := 0
	dp := make([][26]int, len(colors))
	for len(q) != 0 {
		u := q[0]
		q = q[1:]
		cnt++
		dp[u][colors[u]-'a']++
		res = max(res, dp[u][colors[u]-'a'])
		for _, v := range g[u] {
			for i := range dp[u] {
				dp[v][i] = max(dp[v][i], dp[u][i])
			}
			in[v]--
			if in[v] == 0 {
				q = append(q, v)
			}
		}
	}

	if cnt < len(colors) {
		return -1
	}
	return res
}
