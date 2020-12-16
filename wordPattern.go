package main

import "strings"

// Leetcode 290. (easy)
func wordPattern(pattern string, s string) bool {
	strs := strings.Split(s, " ")
	if len(pattern) != len(strs) {
		return false
	}

	m := make(map[byte]string)
	exist := make(map[string]bool)
	for i := range pattern {
		if _, ok := m[pattern[i]]; !ok {
			if exist[strs[i]] {
				return false
			}
			m[pattern[i]] = strs[i]
			exist[strs[i]] = true
			continue
		}
		if m[pattern[i]] != strs[i] {
			return false
		}
	}
	return true
}
