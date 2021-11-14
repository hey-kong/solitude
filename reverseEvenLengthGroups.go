package main

// Leetcode 5927. (medium)
func reverseEvenLengthGroups(head *ListNode) *ListNode {
	i := 0
	var prevLast *ListNode = nil
	prev, cur := prevLast, head
	for cur != nil {
		i++
		j := 0
		for j < i {
			if cur == nil {
				break
			}
			prev = cur
			cur = cur.Next
			j++
		}
		if j > 0 && j%2 == 0 {
			prevLast, prevLast.Next = prevLast.Next, subReverseEvenLengthGroups(prevLast.Next, cur)
		} else {
			prevLast = prev
		}
	}
	return head
}

func subReverseEvenLengthGroups(head, tail *ListNode) *ListNode {
	prev, cur := tail, head
	for cur != tail {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}
	return prev
}
