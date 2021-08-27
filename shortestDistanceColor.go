package main

// Leetcode 1182. (medium)
func shortestDistanceColor(colors []int, queries [][]int) (res []int) {
	n := len(colors)
	left, right := make([][3]int, n), make([][3]int, n)
	for i := 0; i < n; i++ {
		left[i] = [3]int{-1, -1, -1}
		right[i] = [3]int{-1, -1, -1}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			if i-1 >= 0 && left[i-1][j] != -1 {
				left[i][j] = left[i-1][j] + 1
			}
		}
		left[i][colors[i]-1] = 0
	}

	for i := n - 1; i >= 0; i-- {
		for j := 0; j < 3; j++ {
			if i+1 < n && right[i+1][j] != -1 {
				right[i][j] = right[i+1][j] + 1
			}
		}
		right[i][colors[i]-1] = 0
	}

	for _, query := range queries {
		idx, color := query[0], query[1]-1
		d := 1<<31 - 1
		if left[idx][color] != -1 {
			d = left[idx][color]
		}
		if right[idx][color] != -1 {
			d = min(d, right[idx][color])
		}
		if d == 1<<31-1 {
			res = append(res, -1)
		} else {
			res = append(res, d)
		}
	}
	return
}
