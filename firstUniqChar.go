package main

// Leetcode 387. (easy)
func firstUniqChar(s string) int {
	cnt := make([]int, 26)
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}
