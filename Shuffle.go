package main

import "math/rand"

// Leetcode 384. (medium)
type Solution struct {
	nums []int
	shuf []int
}

func Constructor(nums []int) Solution {
	shuf := make([]int, len(nums))
	copy(shuf, nums)
	return Solution{
		nums: nums,
		shuf: shuf,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	for i := len(this.shuf) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		this.shuf[i], this.shuf[j] = this.shuf[j], this.shuf[i]
	}
	return this.shuf
}
