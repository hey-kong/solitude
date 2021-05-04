package main

// Leetcode 206. (easy)
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
