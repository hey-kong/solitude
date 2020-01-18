package main

// Leetcode 76. (hard)
func minWindow(s string, t string) string {
	m := make(map[byte]int)
	for i := range t {
		m[t[i]]++
	}
	lt := len(t)
	start, end := 0, len(s)
	i, j := 0, 0
	for {
		for i < len(s) {
			if _, ok := m[s[i]]; ok {
				m[s[i]]--
				if m[s[i]] >= 0 {
					lt--
				}
				if lt == 0 {
					break
				}
			}
			i++
		}
        if i == len(s) {
			break
		}
		for j < len(s) {
			if _, ok := m[s[j]]; ok {
				m[s[j]]++
				if m[s[j]] > 0 {
					lt++
					break
				}
			}
			j++
		}
		if end - start >= i - j {
			start, end = j, i
		}
		i++
		j++
	}
	if end - start == len(s) {
		return ""
	}
	return s[start:end+1]
}
