package main

import "math/rand"

// Leetcode 398. (medium)
func (this *Solution) Pick(target int) int {
	res := -1
	cnt := 0
	for i, v := range this.nums {
		if v == target {
			cnt++
			// `rand.Intn(cnt)` generate a random number ranging [0,cnt)
			// the probability of `rand.Intn(cnt) == 0` is 1/cnt
			if rand.Intn(cnt) == 0 {
				res = i
			}
		}
	}
	return res
}

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{
		nums: nums,
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */