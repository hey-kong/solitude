package main

// Leetcode 25. (hard)
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	tail := dummy
	num := 0
	for head != nil {
		num++
		if num == k {
			num = 0
			next := head.Next
			tail.Next, head = reverseSublist(tail.Next, head)
			tail = head
			head.Next = next
		}
		head = head.Next
	}
	return dummy.Next
}

func reverseSublist(head, tail *ListNode) (*ListNode, *ListNode) {
	tail.Next = nil
	var pre, cur *ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return tail, head
}
