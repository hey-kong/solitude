package main

// Leetcode 763. (medium)
func partitionLabels(S string) []int {
	m := make(map[byte]int)
	for i := range S {
		m[S[i]] = i
	}
	start, end := 0, 0
	res := []int{}
	for i := range S {
		end = max(end, m[S[i]])
		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}
	return res
}
