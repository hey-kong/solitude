package main

import "sort"

type IntArr []int

// Leetcode 253. (medium)
func minMeetingRooms(intervals [][]int) int {
	start := make([]int, len(intervals))
	end := make([]int, len(intervals))
	for i, interval := range intervals {
		start[i] = interval[0]
		end[i] = interval[1]
	}
	sort.Sort(IntArr(start))
	sort.Sort(IntArr(end))
	room := 0
	i, j := 0, 0
	for i < len(start) {
		if start[i] >= end[j] {
			room = max(room, i-j)
			for end[j] <= start[i] {
				j++
			}
		}
		i++
	}
	return max(room, i-j)
}

func (arr IntArr) Len() int {
	return len(arr)
}

func (arr IntArr) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr IntArr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
