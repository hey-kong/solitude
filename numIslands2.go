package main

// Leetcode 305. (hard)
func numIslands2(m int, n int, positions [][]int) (res []int) {
	root := make([]int, m*n)
	for i := range root {
		root[i] = i
	}

	direction := [4][2]int{[2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}, [2]int{-1, 0}}
	delta := [2]int{n, 1}
	cnt := 0
	visited := make([]bool, m*n)
	for _, position := range positions {
		x, y := position[0], position[1]
		index := x*n + y
		if visited[index] {
			res = append(res, cnt)
			continue
		}
		visited[index] = true
		cnt++
		for i := range direction {
			newX, newY := x+direction[i][0], y+direction[i][1]
			newIndex := index + direction[i][0]*delta[0] + direction[i][1]*delta[1]
			if newX >= 0 && newX < m && newY >= 0 && newY < n &&
				visited[newIndex] && !isConnectedOf305(index, newIndex, root) {
				cnt = unionOf305(index, newIndex, root, cnt)
			}
		}
		res = append(res, cnt)
	}
	return res
}

func findOf305(x int, root []int) int {
	for x != root[x] {
		x = root[x]
	}
	return x
}

func isConnectedOf305(x, y int, root []int) bool {
	return findOf305(x, root) == findOf305(y, root)
}

func unionOf305(x, y int, root []int, cnt int) int {
	x, y = findOf305(x, root), findOf305(y, root)
	if x == y {
		return cnt
	}
	root[x] = y
	cnt--
	return cnt
}
