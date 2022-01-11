package main

import "strings"

// Leetcode 395. (medium)
func longestSubstring(s string, k int) (res int) {
	if s == "" {
		return
	}

	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}

	split := byte(0)
	for i, c := range cnt {
		if c > 0 && c < k {
			split = byte('a' + i)
			break
		}
	}

	if split == 0 {
		return len(s)
	}

	for _, subStr := range strings.Split(s, string(split)) {
		res = max(res, longestSubstring(subStr, k))
	}
	return res
}
