package main

// Leetcode 743. (medium)
func networkDelayTime(times [][]int, n int, k int) (res int) {
	// Init graph
	graph := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		graph[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			graph[i][j] = 0x3f3f3f3f
			if i == j {
				graph[i][j] = 0
			}
		}
	}
	for _, t := range times {
		graph[t[0]][t[1]] = t[2]
	}

	// Dijkstra
	v := make([]bool, n+1)
	for j := 0; j < n; j++ {
		tmp := 0
		for i := 1; i <= n; i++ {
			if !v[i] && (tmp == 0 || graph[k][i] < graph[k][tmp]) {
				tmp = i
			}
		}

		v[tmp] = true

		for i := 1; i <= n; i++ {
			graph[k][i] = min(graph[k][i], graph[k][tmp]+graph[tmp][i])
		}
	}

	for i := 1; i <= n; i++ {
		res = max(res, graph[k][i])
	}
	if res == 0x3f3f3f3f {
		return -1
	}
	return res
}
