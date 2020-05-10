package main

// Leetcode 5404. (easy)
func buildArray(target []int, n int) []string {
	strs := []string{"Push", "Pop"}
	res := []string{}
	idx := 0
	for i := 1; i <= n; i++ {
		if i == target[idx] {
			res = append(res, strs[0])
			idx++
			if idx == len(target) {
				break
			}
		} else {
			res = append(res, strs...)
		}
	}
	return res
}
