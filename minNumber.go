package main

import (
	"sort"
	"strconv"
)

type IntStr []int

// Leetcode m45. (medium)
func minNumber(nums []int) string {
	sort.Sort(IntStr(nums))
	res := ""
	for _, num := range nums {
		res += strconv.Itoa(num)
	}
	return res
}

func (arr IntStr) Len() int {
	return len(arr)
}

func (arr IntStr) Less(i, j int) bool {
	s1 := strconv.Itoa(arr[i]) + strconv.Itoa(arr[j])
	s2 := strconv.Itoa(arr[j]) + strconv.Itoa(arr[i])
	return s1 < s2
}

func (arr IntStr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
