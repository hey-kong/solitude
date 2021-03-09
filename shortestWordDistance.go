package main

// Leetcode 245. (medium)
func shortestWordDistance(wordsDict []string, word1 string, word2 string) int {
	res := len(wordsDict)
	w1, w2 := -1, -1
	same := word1 == word2
	for i, word := range wordsDict {
		if same {
			if word == word1 {
				if w1 == -1 {
					w1 = i
				} else {
					w2 = i
				}
			}
			if w1 != -1 && w2 != -1 {
				res = min(res, w2-w1)
				w1, w2 = w2, -1
			}
		} else {
			if word == word1 {
				w1 = i
			} else if word == word2 {
				w2 = i
			}
			if w1 != -1 && w2 != -1 {
				res = min(res, abs(w2-w1))
			}
		}
	}
	return res
}
