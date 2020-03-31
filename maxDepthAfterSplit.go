package main

// Leetcode 1111. (medium)
func maxDepthAfterSplit(seq string) []int {
	l, r := 0, 0
	res := make([]int, len(seq))
	for i := range seq {
		if seq[i] == '(' {
			res[i] = l
			l = (l + 1) % 2
		} else {
			res[i] = r
			r = (r + 1) % 2
		}
	}
	return res
}
