package main

import "strconv"

// Leetcode 149. (hard)
func maxPoints(points [][]int) int {
	n := len(points)
	res := 0
	for i := 0; i < n; i++ {
		m := make(map[string]int)
		cur := 0
		for j := i + 1; j < n; j++ {
			x1, y1, x2, y2 := points[i][0], points[i][1], points[j][0], points[j][1]
			dx, dy := x2-x1, y2-y1
			if dx == 0 {
				m["x"]++
				cur = max(cur, m["x"])
				continue
			} else if dy == 0 {
				m["y"]++
				cur = max(cur, m["y"])
				continue
			}
			if dx < 0 {
				dx, dy = -dx, -dy
			}
			c := gcd(dx, dy)
			k := strconv.Itoa(dy/c) + "/" + strconv.Itoa(dx/c)
			m[k]++
			cur = max(cur, m[k])
		}
		res = max(res, cur)
	}
	return res + 1
}
