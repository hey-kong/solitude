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
