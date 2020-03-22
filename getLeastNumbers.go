package main

import "container/heap"
import "math/rand"

// Leetcode m40. (easy)
func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return nil
	}

	h := make(MaxHeap, k)
	for i := 0; i < k; i++ {
		h[i] = arr[i]
	}
	heap.Init(&h)
	for i := k; i < len(arr); i++ {
		if h[0] > arr[i] {
			heap.Pop(&h)
			heap.Push(&h, arr[i])
		}
	}
	return h
}

func getLeastNumbers2(arr []int, k int) []int {
	r := rand.Intn(len(arr))
	pivot := arr[r]
	arr[r], arr[0] = arr[0], arr[r]
	mark := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			arr[mark], arr[i] = arr[i], arr[mark]
			mark++
		}
	}

	leftCount := mark - 1
	if mark == k {
		return arr[:mark]
	}
	if leftCount == k {
		return arr[1:mark]
	}
	if leftCount > k {
		return getLeastNumbers2(arr[1:mark], k)
	}
	left := arr[:mark]
	right := getLeastNumbers2(arr[mark:], k-mark)
	return append(left, right...)
}
