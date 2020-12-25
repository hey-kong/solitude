package main

import "sort"

// Leetcode 455. (easy)
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			res++
			j++
			i++
		} else {
			j++
		}
	}
	return res
}
