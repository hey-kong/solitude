package main

// Leetcode m33. (medium)
func verifyPostorder(postorder []int) bool {
	return recursiveVerifyPostorder(postorder, 0, len(postorder)-1)
}

func recursiveVerifyPostorder(postorder []int, l, r int) bool {
	if l >= r {
		return true
	}
	p := l
	for postorder[p] < postorder[r] {
		p++
	}
	q := p
	for postorder[q] > postorder[r] {
		q++
	}
	return q == r && recursiveVerifyPostorder(postorder, l, p-1) && recursiveVerifyPostorder(postorder, p, r-1)
}
