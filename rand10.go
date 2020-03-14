package main

// Leetcode 470. (medium)
func rand10() int {
	a, b, tmp := 0, 0, 0
	for {
		a, b = rand7(), rand7()
		tmp = (a-1)*7 + b
		if tmp <= 40 {
			return tmp%10 + 1
		}
		a = tmp - 40
		b = rand7()
		tmp = (a-1)*7 + b
		if tmp <= 60 {
			return tmp%10 + 1
		}
		a = tmp - 60
		b = rand7()
		tmp = (a-1)*7 + b
		if tmp <= 20 {
			return tmp%10 + 1
		}
	}
	return 0
}
