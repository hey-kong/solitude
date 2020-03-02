package main

import (
	"sort"
	"strconv"
)

// Leetcode 179. (medium)
func largestNumber(nums []int) string {
	sort.Sort(IntStr(nums))
	res := ""
	allIsZero := true
	for _, num := range nums {
		if allIsZero && num != 0 {
			allIsZero = false
		}
		res += strconv.Itoa(num)
	}
	if allIsZero {
		return "0"
	}
	return res
}

/*
func (arr IntStr) Len() int {
	return len(arr)
}

func (arr IntStr) Less(i, j int) bool {
	s1 := strconv.Itoa(arr[i]) + strconv.Itoa(arr[j])
	s2 := strconv.Itoa(arr[j]) + strconv.Itoa(arr[i])
	return s1 > s2
}

func (arr IntStr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
*/
