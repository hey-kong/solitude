package main

// Leetcode 1011. (medium)
func shipWithinDays(weights []int, D int) int {
	left, right := 0, 0
	for _, w := range weights {
		left = max(left, w)
		right += w
	}

	for left <= right {
		curWeights := 0
		curDay := 1
		mid := left + (right-left)/2
		for _, w := range weights {
			if curWeights+w > mid {
				curWeights = 0
				curDay++
			}
			curWeights += w
		}
		if curDay <= D {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
