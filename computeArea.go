package main

// Leetcode 223. (medium)
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	area1, area2 := (C-A)*(D-B), (G-E)*(H-F)
	if C <= E || A >= G || B >= H || D <= F {
		return area1 + area2
	}
	width := min(C, G) - max(A, E)
	height := min(D, H) - max(B, F)
	return area1 + area2 - width*height
}
