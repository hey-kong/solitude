package main

import "sort"

// Leetcode 945. (medium)
func minIncrementForUnique(A []int) int {
	res := 0
	sort.Ints(A)
	for i := 1; i < len(A); i++ {
		if A[i-1] >= A[i] {
			delta := A[i-1] - A[i] + 1
			A[i] += delta
			res += delta
		}
	}
	return res
}

func minIncrementForUnique2(A []int) int {
	pos := make([]int, 80000)
	for i := range pos {
		pos[i] = -1
	}
	res := 0
	for _, num := range A {
		val := findPos(pos, num)
		res += val - num
	}
	return res
}

func findPos(pos []int, i int) int {
	if pos[i] == -1 {
		pos[i] = i
		return i
	}
	val := findPos(pos, pos[i]+1)
	pos[i] = val
	return val
}
