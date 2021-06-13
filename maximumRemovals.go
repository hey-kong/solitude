package main

// Leetcode 5786. (medium)
func maximumRemovals(s string, p string, removable []int) int {
	left, right := 0, len(removable)-1
	for left <= right {
		mid := left + (right-left)/2
		m := make(map[int]bool)
		for i := 0; i <= mid; i++ {
			m[removable[i]] = true
		}
		if existsOfMaximumRemovals(s, p, m) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func existsOfMaximumRemovals(s string, p string, remove map[int]bool) bool {
	i, j := 0, 0
	for i < len(s) && j < len(p) {
		if s[i] == p[j] && !remove[i] {
			j++
		}
		i++
	}
	return j == len(p)
}
