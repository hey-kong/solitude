package main

// Leetcode 323. (medium)
func countComponents(n int, edges [][]int) int {
	root, size := make([]int, n), make([]int, n)
	for i := range root {
		root[i] = i
		size[i] = 1
	}

	for _, edge := range edges {
		n = unionOf323(edge[0], edge[1], root, size, n)
	}
	return n
}

func findOf323(x int, root []int) int {
	for x != root[x] {
		x = root[x]
	}
	return x
}

func unionOf323(x, y int, root, size []int, cnt int) int {
	x, y = findOf323(x, root), findOf323(y, root)
	if x == y {
		return cnt
	}

	if size[x] > size[y] {
		root[y] = x
		size[x] += size[y]
	} else {
		root[x] = y
		size[y] += size[x]
	}
	cnt--
	return cnt
}
