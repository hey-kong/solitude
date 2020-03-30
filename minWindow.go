package main

// Leetcode 76. (hard)
func minWindow(s string, t string) string {
	mapT := make(map[byte]int)
	for i := range t {
		mapT[t[i]]++
	}
	lt := len(t)
	start, end := 0, len(s)
	i, j := 0, 0
	for {
		for j < len(s) {
			if _, ok := mapT[s[j]]; ok {
				mapT[s[j]]--
				if mapT[s[j]] >= 0 {
					lt--
					if lt == 0 {
						break
					}
				}
			}
			j++
		}
		if j == len(s) {
			break
		}
		for i <= j {
			if _, ok := mapT[s[i]]; ok {
				mapT[s[i]]++
				if mapT[s[i]] > 0 {
					lt++
					if j-i < end-start {
						start, end = i, j
					}
					break
				}
			}
			i++
		}
		i++
		j++
	}
	if end-start == len(s) {
		return ""
	}
	return s[start : end+1]
}
