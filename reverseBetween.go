package main

// Leetcode 92. (medium)
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var start *ListNode = nil
	prev := head
	for i := 1; i < m; i++ {
		start = prev
		prev = prev.Next
	}

	export := prev
	curr := prev.Next
	for j := m; j < n; j++ {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// m = 1
	if start == nil {
		export.Next = curr
		return prev
	}
	// m > 1
	start.Next = prev
	export.Next = curr
	return head
}
