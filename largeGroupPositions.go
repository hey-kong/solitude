package main

// Leetcode 830. (easy)
func largeGroupPositions(s string) (res [][]int) {
	cnt := 1
	for i := range s {
		if i == len(s)-1 || s[i] != s[i+1] {
			if cnt >= 3 {
				res = append(res, []int{i - cnt + 1, i})
			}
			cnt = 1
		} else {
			cnt++
		}
	}
	return
}
