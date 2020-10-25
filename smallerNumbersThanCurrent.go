package main

// Leetcode 1365. (easy)
func smallerNumbersThanCurrent(nums []int) []int {
	arr := make([]int, 101)
	for _, num := range nums {
		arr[num]++
	}
	for i := 1; i < 101; i++ {
		arr[i] += arr[i-1]
	}
	for i, num := range nums {
		if num == 0 {
			nums[i] = 0
		} else {
			nums[i] = arr[num-1]
		}
	}
	return nums
}
