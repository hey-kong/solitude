package main

// Leetcode 8. (medium)
func myAtoi(str string) int {
	i := 0
	flag := 1
	var res int64 = 0
	for i < len(str) && str[i] == ' ' {
		i++
	}
	if i == len(str) {
		return 0
	}
	if str[i] == '-' {
		flag = -1
		i++
	} else if str[i] == '+' {
		i++
	}
	for i < len(str) && str[i] >= '0' && str[i] <= '9' {
		res = res*10 + int64(str[i]-'0')
		if flag == 1 && res >= 1<<31-1 {
			return 1<<31 - 1
		} else if flag == -1 && -res <= -1<<31 {
			return -1 << 31
		}
		i++
	}
	return int(res) * flag
}
