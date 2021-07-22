package main

// Leetcode 1711. (medium)
func countPairs(deliciousness []int) (res int) {
	maxVal := deliciousness[0]
	for _, val := range deliciousness[1:] {
		maxVal = max(maxVal, val)
	}

	maxSum := maxVal * 2
	m := make(map[int]int)
	for _, val := range deliciousness {
		for sum := 1; sum <= maxSum; sum <<= 1 {
			res += m[sum-val]
		}
		m[val]++
	}
	return res % (1e9 + 7)
}
