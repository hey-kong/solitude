package main

import "sort"

type Intervals [][]int

// Leetcode 252. (medium)
func canAttendMeetings(intervals [][]int) bool {
	if len(intervals) <= 1 {
		return true
	}
	sort.Sort(Intervals(intervals))
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}

func (intervals Intervals) Len() int {
	return len(intervals)
}

func (intervals Intervals) Less(i, j int) bool {
	return intervals[i][0] < intervals[j][0]
}

func (intervals Intervals) Swap(i, j int) {
	intervals[i], intervals[j] = intervals[j], intervals[i]
}
