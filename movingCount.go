package main

// Leetcode m13. (medium)
func movingCount(m int, n int, k int) int {
	res := 0
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	queue := [][4]int{[4]int{0, 0, 0, 0}}
	for len(queue) > 0 {
		arr := queue[0]
		queue = queue[1:]
		x, y, sumX, sumY := arr[0], arr[1], arr[2], arr[3]
		if x >= m || y >= n || sumX+sumY > k || visited[x][y] {
			continue
		}
		visited[x][y] = true
		res++
		nextX := 0
		if (x+1)%10 == 0 {
			nextX = sumX - 8
		} else {
			nextX = sumX + 1
		}
		nextY := 0
		if (y+1)%10 == 0 {
			nextY = sumY - 8
		} else {
			nextY = sumY + 1
		}
		queue = append(queue, [4]int{x + 1, y, nextX, sumY})
		queue = append(queue, [4]int{x, y + 1, sumX, nextY})
	}
	return res
}
