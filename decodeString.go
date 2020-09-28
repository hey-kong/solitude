package main

import "strings"

// Leetcode 394. (medium)
func decodeString(s string) string {
	res := ""
	num := 0
	strStack := []string{}
	numStack := []int{}
	for i, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			res += s[i : i+1]
		}
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		}
		if c == '[' {
			strStack = append(strStack, res)
			numStack = append(numStack, num)
			res = ""
			num = 0
		}
		if c == ']' {
			res = strStack[len(strStack)-1] + strings.Repeat(res, numStack[len(numStack)-1])
			strStack = strStack[:len(strStack)-1]
			numStack = numStack[:len(numStack)-1]
		}
	}
	return res
}
