package main

import "github.com/sqs/goreturns/returns"

// Leetcode 1. (easy)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	res := make([]int, 2)
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok {
			res[0] = i
			res[1] = j
			break
		}
		m[num] = i
	}
	return res
}

// Leetcode 167. (easy)
func twoSum2(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}
	return []int{-1, -1}
}
