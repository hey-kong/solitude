package main

// Leetcode 466. (hard)
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	cnt1, cnt2 := 0, 0
	j := 0
	for cnt1 < n1 {
		for i := 0; i < len(s1); i++ {
			if s1[i] == s2[j] {
				if j == len(s2)-1 {
					j = 0
					cnt2++
				} else {
					j++
				}
			}
		}
		cnt1++
	}
	return cnt2 / n2
}
