package main

// Leetcode 148. (medium)
func sortList(head *ListNode) *ListNode {
	return mergeSortList(head, nil)
}

func mergeSortList(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	return mergeTwoLists(mergeSortList(head, slow), mergeSortList(slow, tail))
}

func quickSortList(head, tail *ListNode) *ListNode {
	if head == tail || head.Next == tail {
		return head
	}

	mid := partitionSortList(head, tail)
	quickSortList(head, mid)
	quickSortList(mid.Next, tail)
	return head
}

func partitionSortList(head, tail *ListNode) *ListNode {
	pivot := head.Val
	pre, mark := head, head.Next
	cur := head.Next
	for cur != tail {
		if cur.Val < pivot {
			mark.Val, cur.Val = cur.Val, mark.Val
			pre, mark = mark, mark.Next
		}
		cur = cur.Next
	}
	head.Val, pre.Val = pre.Val, head.Val
	return pre
}
