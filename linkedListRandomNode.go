package main

import "math/rand"

// Leetcode 382. (medium)
/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	head := this.head

	var res int
	cnt := 0
	for head != nil {
		cnt++
		// the probability of `rand.Intn(cnt) == 0` is 1/cnt
		if rand.Intn(cnt) == 0 {
			res = head.Val
		}
		head = head.Next
	}
	return res
}

type Solution struct {
	head *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{
		head: head,
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */