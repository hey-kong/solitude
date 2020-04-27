package main

import "container/heap"

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
		l = (l + 1) / 2
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
			<-notify
		}
		l = (l + 1) / 2
	}
	return lists[0]
}

func mergeKLists2(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}

	var h MinNodeHeap
	heap.Init(&h)
	for i := 0; i < l; i++ {
		if lists[i] != nil {
			heap.Push(&h, lists[i])
		}
	}

	dummy := &ListNode{}
	p := dummy
	for h.Len() != 0 {
		tmp := heap.Pop(&h).(*ListNode)
		p.Next = tmp
		p = p.Next
		if tmp.Next != nil {
			heap.Push(&h, tmp.Next)
		}
	}
	return dummy.Next
}

type MinNodeHeap []*ListNode

func (h MinNodeHeap) Len() int {
	return len(h)
}

func (h MinNodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h MinNodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinNodeHeap) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return tmp
}
