package main

// Leetcode 28. (easy)
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	next := getNext(needle)
	i, j := 0, 0
	for i < len(haystack) && j < len(needle) {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(needle) {
		return i - j
	}
	return -1
}
