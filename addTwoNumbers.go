package main

// Leetcode 445. (medium)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	length1, length2 := 0, 0
	h1, h2 := l1, l2
	for h1 != nil {
		h1 = h1.Next
		length1++
	}
	for h2 != nil {
		h2 = h2.Next
		length2++
	}

	if length1 < length2 {
		l1, l2 = l2, l1
		length1, length2 = length2, length1
	}
	for i := 0; i < length1 - length2; i++ {
		pre := &ListNode{
			Val:  0,
			Next: l2,
		}
		l2 = pre
	}

	if offset := recursiveAddTwoNumbers(l1, l2); offset {
		pre := &ListNode{
			Val:  1,
			Next: l1,
		}
		l1 = pre
	}
	return l1
}

func recursiveAddTwoNumbers(l1 *ListNode, l2 *ListNode) bool {
	if l1 == nil {
		return false
	}

	offset := recursiveAddTwoNumbers(l1.Next, l2.Next)
	if offset {
		l1.Val += l2.Val + 1
	} else {
		l1.Val += l2.Val
	}

	if l1.Val >= 10 {
		l1.Val -= 10
		return true
	}
	return false
}
