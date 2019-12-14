package main

// Leetcode 23. (hard)
func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}

	for l > 1 {
		for i := 0; i < l/2; i++ {
			lists[i] = mergeTwoLists(lists[i], lists[l-1-i])
		}
		l = (l+1) / 2
	}
	return lists[0]
}

func mergeKListsWithGoroutine(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}

	notify := make(chan bool)
	for l > 1 {
		for i := 0; i < l/2; i++ {
			go func() {
				lists[i] = mergeTwoLists(lists[i], lists[l-1-i])
				notify <- true
			}()
		}
		for i := 0; i < l/2; i++ {
			<- notify
		}
		l = (l+1) / 2
	}
	return lists[0]
}
