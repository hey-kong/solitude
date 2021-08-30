package main

// Leetcode 1130. (medium)
func mctFromLeafValues(arr []int) int {
	res := 0
	stack := []int{1<<31}
	for i := 0; i < len(arr); i++ {
		for stack[len(stack)-1] < arr[i] {
			res += stack[len(stack)-1] * min(stack[len(stack)-2], arr[i])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, arr[i])
	}
	for len(stack) > 2 {
		res += stack[len(stack)-1] * stack[len(stack)-2]
		stack = stack[:len(stack)-1]
	}
	return res
}
