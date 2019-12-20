package main

// Leetcode 19. (medium)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cnt := 0
	prev := head
	for prev != nil {
		cnt++
		prev = prev.Next
	}

	prev = nil
	curr := head
	k := cnt - n
	for k > 0 {
		prev = curr
		curr = curr.Next
		k--
	}

	if prev == nil {
		return curr.Next
	}

	prev.Next = curr.Next
	return head
}
