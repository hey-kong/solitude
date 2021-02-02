package main

// Leetcode 30. (hard)
func findSubstring(s string, words []string) (res []int) {
	m := make(map[string]int)
	for _, w := range words {
		m[w]++
	}
	l, n := len(words[0]), len(words)

	for i := 0; i < l; i++ {
		left, right := i, i
		tmpMap := make(map[string]int)
		cnt := 0
		for right+l <= len(s) {
			w := s[right : right+l]
			right += l
			if m[w] == 0 {
				left = right
				cnt = 0
				tmpMap = map[string]int{}
			} else {
				tmpMap[w]++
				cnt++
				for tmpMap[w] > m[w] {
					tmpMap[s[left:left+l]]--
					left += l
					cnt--
				}
				if cnt == n {
					res = append(res, left)
				}
			}
		}

	}
	return res
}
