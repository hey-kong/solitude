package main

// Leetcode 127. (medium)
func ladderLength(beginWord string, endWord string, wordList []string) int {
	m := make(map[string]bool)
	for _, s := range wordList {
		m[s] = true
	}
	queue := make([]string, 1)
	queue[0] = beginWord
	level := 1
	for len(queue) != 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			word := queue[0]
			queue = queue[1:]
			if word == endWord {
				return level
			}
			for j := 0; j < len(word); j++ {
				for b := 'a'; b <= 'z'; b++ {
					tmpWord := word[:j] + string(b) + word[j+1:]
					if m[tmpWord] {
						m[tmpWord] = false
						queue = append(queue, tmpWord)
					}
				}
			}
		}
		level++
	}
	return 0
}
