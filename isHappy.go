package main

// Leetcode 202. (easy)
func isHappy(n int) bool {
	slow, fast := n, n
	for {
		slow = squareSumOfIsHappy(slow)
		fast = squareSumOfIsHappy(squareSumOfIsHappy(fast))
		if slow == fast {
			break
		}
	}
	return slow == 1
}

func squareSumOfIsHappy(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
