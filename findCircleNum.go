package main

// Leetcode 547. (medium)
func findCircleNum(M [][]int) int {
	n := len(M)
	root, size := make([]int, n), make([]int, n)
	for i := range root {
		root[i] = i
	}
	for i := range size {
		size[i] = 1
	}
	cnt := n
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if M[i][j] == 1 {
				cnt, root, size = union(i, j, cnt, root, size)
			}
		}
	}
	return cnt
}

func find(x int, root []int) int {
	for x != root[x] {
		x = root[x]
	}
	return x
}

func union(a, b, cnt int, root, size []int) (int, []int, []int) {
	x := find(a, root)
	y := find(b, root)
	if x == y {
		return cnt, root, size
	}
	if size[x] < size[y] {
		root[x] = y
		size[y] += size[x]
	} else {
		root[y] = x
		size[x] += size[y]
	}
	cnt--
	return cnt, root, size
}
