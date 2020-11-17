package main

import "strconv"

// Leetcode 43. (medium)
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	n := ""
	nums := make([]string, len(num2))
	for i := len(num2) - 1; i >= 0; i-- {
		nums[i] = strMultiplyInt(num1, int(num2[i]-'0')) + n
		n += "0"
	}
	for i := 1; i < len(num2); i++ {
		nums[0] = addStrings(nums[0], nums[i])
	}
	return nums[0]
}

func strMultiplyInt(str string, num int) string {
	res := ""
	flag := 0
	for i := len(str) - 1; i >= 0; i-- {
		n := int(str[i]-'0')*num + flag
		res = strconv.Itoa(n%10) + res
		flag = n / 10
	}
	if flag != 0 {
		res = strconv.Itoa(flag) + res
	}
	return res
}
