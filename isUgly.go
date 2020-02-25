package main

// Leetcode 263. (easy)
func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	for num != 1 {
		tmp := num
		if num%2 == 0 {
			num /= 2
		}
		if num%3 == 0 {
			num /= 3
		}
		if num%5 == 0 {
			num /= 5
		}
		if num == tmp {
			return false
		}
	}
	return true
}
