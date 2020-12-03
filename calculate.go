package main

// Leetcode m16.26. (medium)
func calculate(s string) int {
	stack := []int{}
	i := 0
	for i < len(s) {
		if s[i] == ' ' {
			i++
			continue
		}
		op := s[i]
		if op == '*' || op == '/' || op == '+' || op == '-' {
			i++
			for i < len(s) && s[i] == ' ' {
				i++
			}
		}
		num := 0
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
			i++
		}
		if op == '-' {
			num = -num
		}
		if op == '*' {
			num = stack[len(stack)-1] * num
			stack = stack[:len(stack)-1]
		}
		if op == '/' {
			num = stack[len(stack)-1] / num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}
	res := 0
	for _, num := range stack {
		res += num
	}
	return res
}
