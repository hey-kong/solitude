package main

// Leetcode 5702. (medium)
func findCenter(edges [][]int) int {
	n := 0
	for i := range edges {
		if n < edges[i][0] {
			n = edges[i][0]
		}
		if n < edges[i][1] {
			n = edges[i][1]
		}
	}

	d := make([]int, n)
	for i := range edges {
		d[edges[i][0]-1]++
		if d[edges[i][0]-1] == n-1 {
			return edges[i][0]
		}
		d[edges[i][1]-1]++
		if d[edges[i][1]-1] == n-1 {
			return edges[i][1]
		}
	}
	return 0
}
