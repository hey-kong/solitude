package main

// Leetcode 159. (medium)
func lengthOfLongestSubstringTwoDistinct(s string) int {
	n := len(s)
	if n < 3 {
		return n
	}

	m := make(map[byte]int, 3)
	res := 2
	left, right := 0, 0
	for right < n {
		m[s[right]] = right
		if len(m) == 3 {
			var preByte byte
			preIdx := n
			for k, v := range m {
				if v < preIdx {
					preByte = k
					preIdx = v
				}
			}
			delete(m, preByte)
			res = max(res, right-left)
			left = preIdx + 1
		}
		right++
	}
	res = max(res, right-left)
	return res
}
