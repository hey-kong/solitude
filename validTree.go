package main

// Leetcode 261. (medium)
func validTree(n int, edges [][]int) bool {
	if len(edges) < n-1 {
		return false
	}

	degree := make([]int, n)
	visited := make([]bool, n)
	for _, e := range edges {
		degree[e[0]]++
		degree[e[1]]++
	}
	queue := make([]int, 0)
	for i, d := range degree {
		if d == 0 {
			return n == 1
		} else if d == 1 {
			queue = append(queue, i)
			degree[i] = 0
			visited[i] = true
			n--
		}
	}
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, e := range edges {
			if cur == e[0] && !visited[e[1]] {
				degree[e[1]]--
			} else if cur == e[1] && !visited[e[0]] {
				degree[e[0]]--
			}
		}
		for i, d := range degree {
			if d == 1 {
				queue = append(queue, i)
				degree[i] = 0
				visited[i] = true
				n--
			}
		}
	}
	return n == 0
}
