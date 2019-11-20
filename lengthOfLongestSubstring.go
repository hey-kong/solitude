package main

// Leetcode 3. (medium)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	maxLength := 0
	left := 0
	m := map[rune]int{}
	for i, ch := range s {
		if j, ok := m[ch]; ok {
			left = max(left, j+1)
		}
		m[ch] = i
		maxLength = max(maxLength, i-left+1)
	}
	return maxLength
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
