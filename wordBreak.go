package main

import "strings"

// Leetcode 139. (medium)
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	dict := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		dict[word] = true
	}
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if _, ok := dict[s[j:i]]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// Leetcode 140. (hard)
func wordBreak2(s string, wordDict []string) []string {
	m := make(map[string]bool)
	for _, word := range wordDict {
		m[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= i; j++ {
			if m[s[j-1:i]] && dp[j-1] {
				dp[i] = true
				break
			}
		}
	}

	if !dp[len(s)] {
		return nil
	}
	return dfs(s, m, 0, []string{}, []string{})
}

func dfs(s string, m map[string]bool, begin int, cur, res []string) []string {
	if begin == len(s) {
		tmp := strings.Join(cur, " ")
		res = append(res, tmp)
		return res
	}

	for i := begin; i < len(s); i++ {
		if !m[s[begin:i+1]] {
			continue
		}
		cur = append(cur, s[begin:i+1])
		res = dfs(s, m, i+1, cur, res)
		cur = cur[:len(cur)-1]
	}
	return res
}
