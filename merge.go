package main

import "sort"

// Leetcode 88. (easy)
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	mark := m + n - 1
	for mark >= 0 && i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[mark] = nums2[j]
			j--
		} else {
			nums1[mark] = nums1[i]
			i--
		}
		mark--
	}

	for mark >= 0 && j >= 0 {
		nums1[mark] = nums2[j]
		j--
		mark--
	}
}

// Leetcode 56. (medium)
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Sort(Intervals(intervals))
	j := 0
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] >= intervals[i+1][0] {
			intervals[i+1][0] = intervals[i][0]
			intervals[i+1][1] = max(intervals[i][1], intervals[i+1][1])
		}
		if intervals[j][0] != intervals[i+1][0] {
			j++
		}
		intervals[j] = intervals[i+1]
	}
	return intervals[:j+1]
}
