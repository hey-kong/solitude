package main

// Leetcode 328. (medium)
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	oddHead, evenHead := head, head.Next
	join := head.Next
	for evenHead != nil && evenHead.Next != nil {
		oddHead.Next = evenHead.Next
		oddHead = oddHead.Next
		evenHead.Next = oddHead.Next
		evenHead = evenHead.Next
	}
	oddHead.Next = join
	return head
}
