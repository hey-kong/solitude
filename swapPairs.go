package main

// Leetcode 24. (medium)
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}

func swapPairs2(head *ListNode) *ListNode {
	preHead := &ListNode{0, head}
	pre := preHead
	for pre.Next != nil && pre.Next.Next != nil {
		curr, next := pre.Next, pre.Next.Next
		curr.Next = next.Next
		next.Next = curr
		pre.Next = next
		pre = curr
	}
	return preHead.Next
}
