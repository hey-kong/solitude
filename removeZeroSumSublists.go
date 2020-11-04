package main

// Leetcode 1171. (medium)
func removeZeroSumSublists(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	m := make(map[int]*ListNode)
	sum := 0
	for p := dummy; p != nil; p = p.Next {
		sum += p.Val
		m[sum] = p
	}
	sum = 0
	for p := dummy; p != nil; p = p.Next {
		sum += p.Val
		if node, ok := m[sum]; ok {
			p.Next = node.Next
		}
	}
	return dummy.Next
}
