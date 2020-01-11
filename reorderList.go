package main

// Leetcode 143. (medium)
func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	h1, h2 := head, reverseList(slow.Next)
	slow.Next = nil
	for h1.Next != nil {
		next := h1.Next
		h1.Next = h2
		h1 = next
		next = h2.Next
		h2.Next = h1
		h2 = next
	}
	if h2 != nil {
		h1.Next = h2
	}
}
