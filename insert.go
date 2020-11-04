package main

import "sort"

// Leetcode 57. (hard)
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	if len(intervals) == 1 {
		return intervals
	}

	sort.Sort(Intervals(intervals))
	i := 0
	for j := 1; j < len(intervals); j++ {
		if intervals[i][1] >= intervals[j][0] {
			intervals[i][1] = max(intervals[i][1], intervals[j][1])
		} else {
			i++
			intervals[i] = intervals[j]
		}
	}
	return intervals[:i+1]
}
