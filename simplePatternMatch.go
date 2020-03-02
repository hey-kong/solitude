package main

import "strings"

// match("abba", "Beijing Hangzhou Hangzhou Beijing") -- true
// match("aabb", "Beijing Hangzhou Hangzhou Beijing") -- false
// match("baab", "Beijing Hangzhou Hangzhou Beijing") -- true
func simplePatternMatch(pattern, s string) bool {
	m := make(map[byte]string)
	strs := strings.Split(s, " ")
	if len(pattern) != len(strs) {
		return false
	}
	for i := range pattern {
		if str, ok := m[pattern[i]]; ok {
			if str != strs[i] {
				return false
			}
		} else {
			m[pattern[i]] = strs[i]
		}
	}
	return true
}
