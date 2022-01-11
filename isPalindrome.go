package main

import "strings"

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

// Leetcode 125. (easy)
func isPalindrome2(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i <= j {
		for i <= j && !isValidOfIsPalindrome(s[i]) {
			i++
		}
		for i <= j && !isValidOfIsPalindrome(s[j]) {
			j--
		}
		if i <= j && s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isValidOfIsPalindrome(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') || (b >= 'A' && b <= 'Z')
}
