package main

// Leetcode 739. (medium)
func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	stack := make([]int, 1)
	stack[0] = 0
	for i := 1; i < len(T); i++ {
		for len(stack) > 0 && T[stack[len(stack)-1]] < T[i] {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
