package main

// Leetcode 1207. (easy)
func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, num := range arr {
		m[num]++
	}
	times := make(map[int]bool)
	for _, time := range m {
		if times[time] {
			return false
		} else {
			times[time] = true
		}
	}
	return true
}
