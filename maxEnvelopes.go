package main

import "sort"

type Envelopes [][]int

// Leetcode 354. (hard)
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}

	sort.Sort(Envelopes(envelopes))
	dp := make([]int, len(envelopes))
	for i := range dp {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < len(envelopes); i++ {
		for j := 0; j < i; j++ {
			if envelopes[j][0] < envelopes[i][0] && envelopes[j][1] < envelopes[i][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func (envelopes Envelopes) Len() int {
	return len(envelopes)
}

func (envelopes Envelopes) Less(i, j int) bool {
	if envelopes[i][0] == envelopes[j][0] {
		return envelopes[i][1] < envelopes[j][1]
	}
	return envelopes[i][0] < envelopes[j][0]
}

func (envelopes Envelopes) Swap(i, j int) {
	envelopes[i], envelopes[j] = envelopes[j], envelopes[i]
}
