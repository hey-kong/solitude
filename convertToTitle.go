package main

// Leetcode 168. (easy)
func convertToTitle(n int) string {
	res := ""
	for n > 0 {
		n--
		res = string('A' + n % 26) + res
		n /= 26
	}
	return res
}
