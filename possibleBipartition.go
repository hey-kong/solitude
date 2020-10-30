package main

// Leetcode 886. (medium)
func possibleBipartition(N int, dislikes [][]int) bool {
	root := make([]int, 2*N+1)
	for i := range root {
		root[i] = i
	}
	for i := range dislikes {
		x, y := findOf886(dislikes[i][0], root), findOf886(dislikes[i][1], root)
		if x == y {
			return false
		}
		a, b := findOf886(dislikes[i][0]+N, root), findOf886(dislikes[i][1]+N, root)
		root[b] = x
		root[a] = y
	}
	return true
}

func findOf886(x int, root []int) int {
	for x != root[x] {
		x = root[x]
	}
	return x
}
