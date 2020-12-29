package main

import "container/heap"
import "unsafe"

// Leetcode 1046. (easy)
func lastStoneWeight(stones []int) int {
	h := (*MaxHeap)(unsafe.Pointer(&stones))
	heap.Init(h)
	for len(stones) > 1 {
		x, y := heap.Pop(h).(int), heap.Pop(h).(int)
		if x > y {
			heap.Push(h, (x - y))
		}
	}
	if len(stones) == 1 {
		return (*h)[0]
	}
	return 0
}
