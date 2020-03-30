package main

// Leetcode 239. (hard)
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, len(nums)-k+1)
	queue := []int{}
	i, j := 0, 0
	for j < k {
		for len(queue) > 0 && queue[len(queue)-1] < nums[j] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[j])
		j++
	}
	for j < len(nums) {
		res[i] = queue[0]
		if nums[i] == queue[0] {
			queue = queue[1:]
		}
		i++
		for len(queue) > 0 && queue[len(queue)-1] < nums[j] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[j])
		j++
	}
	res[i] = queue[0]
	return res
}
