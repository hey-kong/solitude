package main

// Leetcode 42. (hard)
func trap(height []int) int {
	res := 0
	for i := 0; i < len(height); i++ {
		left := 0
		right := 0
		for j := 0; j < i; j++ {
			left = max(left, height[j])
		}
		for k := len(height)-1; k > i; k-- {
			right = max(right, height[k])
		}
		res += max(0, min(left,right) - height[i])
	}
	return res
}

func trap2(height []int) int {
	res := 0
	left := make([]int, len(height))
	right := make([]int, len(height))
	for i := 1; i < len(height)-1; i++ {
		left[i] = max(left[i-1], height[i-1])
	}
	for i := len(height)-2; i > 0; i-- {
		right[i] = max(right[i+1], height[i+1])
	}

	for i := 1; i < len(height)-1; i++ {
		res += max(min(left[i], right[i]) - height[i], 0)
	}
	return res
}

func trap3(height []int) int {
	if len(height) < 3 {
		return 0
	}

	res := 0
	leftMax := height[0]
	rightMax := height[len(height)-1]
	left := 0
	right := len(height) - 2
	for left <= right {
		if leftMax < rightMax {
			res += max(leftMax - height[left], 0)
			leftMax = max(leftMax, height[left])
			left++
		} else {
			res += max(rightMax - height[right], 0)
			rightMax = max(rightMax, height[right])
			right--
		}
	}
	return res
}
