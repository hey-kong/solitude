package main

import "sort"

// Leetcode 5779. (hard)
func minWastedSpace(packages []int, boxes [][]int) int {
	sort.Ints(packages)
	res := 1<<63 - 1
	for _, box := range boxes {
		sort.Ints(box)
		if box[len(box)-1] < packages[len(packages)-1] {
			continue
		}

		tmp, l := 0, 0
		for _, b := range box {
			r := sort.SearchInts(packages, b+1)
			tmp += (r - l) * b
			l = r
		}
		res = min(res, tmp)
	}

	if res < 1<<63-1 {
		for _, p := range packages {
			res -= p
		}
		return res % (1e9 + 7)
	}
	return -1
}
