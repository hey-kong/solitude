package main

import "sort"

// Leetcode 1229. (medium)
func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	sort.Sort(Intervals(slots1))
	sort.Sort(Intervals(slots2))
	i, j := 0, 0
	for i < len(slots1) && j < len(slots2) {
		s1, e1 := slots1[i][0], slots1[i][1]
		s2, e2 := slots2[j][0], slots2[j][1]
		ms, me := max(s1, s2), min(e1, e2)
		if me-ms >= duration {
			return []int{ms, ms + duration}
		}
		if e1 < e2 {
			i++
		} else {
			j++
		}
	}
	return nil
}
