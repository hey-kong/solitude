package main

// Leetcode 109. (medium)
func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{
			Val: head.Val,
		}
	}
	var pre *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow, fast = slow.Next, fast.Next.Next
	}
	pre.Next = nil
	node := &TreeNode{
		Val: slow.Val,
	}
	node.Left = sortedListToBST(head)
	node.Right = sortedListToBST(slow.Next)
	return node
}
