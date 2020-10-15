package main

// Leetcode 351. (medium)
func numberOfPatterns(m int, n int) int {
	root := make([]int, 10)
	for i := 1; i <= 9; i++ {
		root[i] = i
	}
	unionOf351(1, 3, root)
	unionOf351(1, 7, root)
	unionOf351(1, 9, root)
	unionOf351(2, 8, root)
	unionOf351(4, 6, root)

	res := 0
	mark := make([]bool, 10)
	res = dfsOfNumberOfPatterns(1, 1, m, n, res, mark, root)
	res = dfsOfNumberOfPatterns(2, 1, m, n, res, mark, root)
	res *= 4
	res = dfsOfNumberOfPatterns(5, 1, m, n, res, mark, root)
	return res
}

func dfsOfNumberOfPatterns(cur, k, m, n, res int, mark []bool, root []int) int {
	if k >= m && k <= n {
		res++
	} else if k > n {
		return res
	}

	mark[cur] = true
	for i := 1; i <= 9; i++ {
		if mark[i] {
			continue
		}
		if flag := findOf351(cur, root) == findOf351(i, root); flag && !mark[(cur+i)/2] {
			continue
		}
		res = dfsOfNumberOfPatterns(i, k+1, m, n, res, mark, root)
	}
	mark[cur] = false
	return res
}

func findOf351(x int, root []int) int {
	for x != root[x] {
		x = root[x]
	}
	return x
}

func unionOf351(x, y int, root []int) []int {
	r1, r2 := findOf351(x, root), findOf351(y, root)
	if r1 != r2 {
		root[r1] = r2
	}
	return root
}
