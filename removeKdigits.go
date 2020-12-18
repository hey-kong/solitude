package main

// Leetcode 402. (medium)
func removeKdigits(num string, k int) string {
	res := []byte{}
	for i := range num {
		for len(res) > 0 && res[len(res)-1] > num[i] && k > 0 {
			res = res[:len(res)-1]
			k--
		}
		res = append(res, num[i])
	}
	start, end := 0, len(res)-k
	for start < len(res) && res[start] == '0' {
		start++
	}

	if start >= end {
		return "0"
	}
	return string(res[start:end])
}
