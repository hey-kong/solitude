package main

// Leetcode 5908. (medium)
func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	g := make([][]int, n)
	for i, parent := range parents[1:] {
		g[parent] = append(g[parent], i+1)
	}

	res := 0
	maxScore := 0
	var dfs (func(int) int)
	dfs = func(p int) int {
		size, score := 1, 1
		for _, c := range g[p] {
			ans := dfs(c)
			size += ans
			score *= ans
		}
		if p > 0 {
			score *= n - size
		}

		if score > maxScore {
			maxScore = score
			res = 1
		} else if score == maxScore {
			res++
		}
		return size
	}
	dfs(0)
	return res
}
