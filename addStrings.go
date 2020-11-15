package main

import "strconv"

// Leetcode 415. (easy)
func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	res := ""
	flag := 0
	i, j := len(num1)-1, len(num2)-1
	for j >= 0 {
		num := int(num1[i]-'0') + int(num2[j]-'0') + flag
		res = strconv.Itoa(num%10) + res
		flag = num / 10
		i--
		j--
	}
	for i >= 0 {
		num := int(num1[i]-'0') + flag
		res = strconv.Itoa(num%10) + res
		flag = num / 10
		i--
	}
	if flag == 1 {
		res = "1" + res
	}
	return res
}
