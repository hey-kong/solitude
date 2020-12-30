package main

import "sort"

// Leetcode 435. (medium)
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	save := 1
	right := intervals[0][1]
	for _, tmp := range intervals[1:] {
		if tmp[0] >= right {
			save++
			right = tmp[1]
		}
	}
	return len(intervals) - save
}
