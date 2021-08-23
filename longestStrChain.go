package main

import "sort"

// Leetcode 1048. (medium)
func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	dp := make([]int, len(words))
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(words); i++ {
		for j := 0; j < i; j++ {
			if checkOfLongestStrChain(words[j], words[i]) {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return maxInt(dp...)
}

func checkOfLongestStrChain(word1, word2 string) bool {
	if len(word1)+1 != len(word2) {
		return false
	}

	flag := false
	i, j := 0, 0
	for i < len(word1) && j < len(word2) {
		if word1[i] == word2[j] {
			i++
			j++
		} else {
			if flag {
				return false
			}
			flag = true
			j++
		}
	}
	return true
}
