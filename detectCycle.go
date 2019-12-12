package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Leetcode 142. (medium)
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			break
		}
	}
	slow = head
	for slow != fast {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}
