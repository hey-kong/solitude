package main

import (
	"strconv"
	"strings"
)

// Leetcode 01.06. (easy)
func compressString(S string) string {
	if len(S) <= 1 {
		return S
	}

	var sb strings.Builder
	pre := S[0]
	cnt := 1
	for i := 1; i < len(S); i++ {
		if S[i] == pre {
			cnt++
		} else {
			sb.WriteByte(pre)
			sb.WriteString(strconv.Itoa(cnt))
			pre = S[i]
			cnt = 1
		}
	}
	sb.WriteByte(pre)
	sb.WriteString(strconv.Itoa(cnt))
	if sb.Len() >= len(S) {
		return S
	}
	return sb.String()
}
