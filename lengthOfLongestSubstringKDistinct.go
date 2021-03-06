package main

// Leetcode 340. (medium)
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	i, j := 0, 0
	m := make(map[byte]int)
	res := 0
	for j < len(s) {
		m[s[j]]++
		if m[s[j]] == 1 {
			k--
			for k < 0 {
				m[s[i]]--
				if m[s[i]] == 0 {
					k++
				}
				i++
			}
		}
		res = max(res, j-i+1)
		j++
	}
	return res
}
