package main

func kmp(s string, t string) int {
	next := getNext(t)
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if j == -1 || s[i] == t[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(t) {
		return i - j
	}
	return -1
}

func getNext(s string) []int {
	next := make([]int, len(s))
	i, j := 0, -1
	next[i] = -1
	for i < len(s)-1 {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}
