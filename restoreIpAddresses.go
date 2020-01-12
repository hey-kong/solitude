package main

import "strings"

// Leetcode 93. (medium)
func restoreIpAddresses(s string) []string {
	res := []string{}
	ip := ""
	res = recursiveRestoreIpAddresses(s, ip, 1, 0, res)
	return res
}

func recursiveRestoreIpAddresses(s, ip string, step, start int, res []string) []string {
	if start == len(s) && step == 5 {
		ip = strings.Trim(ip, ".")
		res = append(res, ip)
		return res
	}
	if len(s) - start < 5 - step {
		return res
	}
	if len(s) - start > 3 * (5 - step) {
		return res
	}

	num := 0
	for i := start; i < start + 3 && i < len(s); i++ {
		num = num * 10 + int(s[i] - '0')
		if num > 255 {
			continue
		}
		ip += s[i:i+1]
		res = recursiveRestoreIpAddresses(s, ip+".", step+1, i+1, res)
		if num == 0 {
			break
		}
	}
	return res
}
