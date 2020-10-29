package main

import "strconv"

// Leetcode 670. (medium)
func maximumSwap(num int) int {
	nums := []byte(strconv.Itoa(num))
	arr := make([]int, 10)
	for i := range nums {
		arr[nums[i]-'0'] = i
	}
	for i := range nums {
		for j := 9; j > int(nums[i]-'0'); j-- {
			if arr[j] > i {
				nums[i], nums[arr[j]] = nums[arr[j]], nums[i]
				res, _ := strconv.Atoi(string(nums))
				return res
			}
		}
	}
	return num
}
