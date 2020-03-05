package main

// Leetcode m57 - II. (easy)
func findContinuousSequence(target int) [][]int {
	l, r := 1, 2
	res := [][]int{}
	for l < r {
		value := (l + r) * (r - l + 1) / 2
		if value == target {
			tmp := make([]int, r-l+1)
			for i := l; i <= r; i++ {
				tmp[i-l] = i
			}
			res = append(res, tmp)
			l++
		} else if value < target {
			r++
		} else if value > target {
			l++
		}
	}
	return res
}
