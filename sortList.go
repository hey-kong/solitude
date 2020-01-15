package main

// Leetcode 148. (medium)
func sortList(head *ListNode) *ListNode {
	preHead := new(ListNode)
	preHead.Next = head
	tmpNode := head
	l := 0
	for tmpNode != nil {
		l++
		tmpNode = tmpNode.Next
	}
	for size := 1; size < l; size <<= 1 {
		pre := preHead
		cur := preHead.Next
		for cur != nil {
			left := cur
			right := sortListCut(left, size)
			cur = sortListCut(right, size)

			pre.Next = sortListMerge(left, right)
			for pre.Next != nil {
				pre = pre.Next
			}
		}
	}
	return preHead.Next
}

func sortListCut(head *ListNode, size int) *ListNode {
	var pre *ListNode
	cur := head
	for size > 0 && cur != nil {
		pre = cur
		cur = cur.Next
		size--
	}
	if pre != nil {
		pre.Next = nil
	}
	return cur
}

func sortListMerge(l1, l2 *ListNode) *ListNode {
	preHead := new(ListNode)
	p := preHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 == nil {
		p.Next = l2
	} else if l2 == nil {
		p.Next = l1
	}
	return preHead.Next
}
