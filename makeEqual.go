package main

// Leetcode 5784. (easy)
func makeEqual(words []string) bool {
	m := make([]int, 26)
	l := len(words)
	for _, word := range words {
		for i := range word {
			m[word[i]-'a']++
		}
	}

	for _, cnt := range m {
		if cnt%l != 0 {
			return false
		}
	}
	return true
}
