package main

import "math/rand"

// Leetcode 528. (medium)
type Solution struct {
	sum []int
}

func Constructor(w []int) Solution {
	s := Solution{
		sum: make([]int, len(w)),
	}
	tmp := 0
	for i := range w {
		tmp += w[i]
		s.sum[i] = tmp
	}
	return s
}

func (this *Solution) PickIndex() int {
	left, right := 0, len(this.sum)-1
	r := rand.Intn(this.sum[right])
	r++
	for left <= right {
		mid := left + (right-left)/2
		if r == this.sum[mid] {
			left = mid
			break
		} else if r < this.sum[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
