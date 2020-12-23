package main

// Leetcode 61. (medium)
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	cnt := 1
	for p.Next != nil {
		p = p.Next
		cnt++
	}
	p.Next = head
	cnt -= k % cnt
	for cnt != 0 {
		p = p.Next
		cnt--
	}
	head = p.Next
	p.Next = nil
	return head
}
