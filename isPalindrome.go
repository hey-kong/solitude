package main

// Leetcode 234. (easy)
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	var pre *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next := slow.Next
		slow.Next = pre
		pre = slow
		slow = next
	}
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		if slow.Val != pre.Val {
			return false
		}
		slow = slow.Next
		pre = pre.Next
	}
	return true
}
