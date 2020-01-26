package main

// Leetcode 222. (medium)
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	d := computeDepth(root)
	if d == 0 {
		return 1
	}

	cnt := 1<<uint(d) - 1
	left, right := 0, 1<<uint(d)-1
	for left <= right {
		mid := left + (right-left)/2
		if existsNode(mid, d, root) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left + cnt
}

func computeDepth(root *TreeNode) int {
	d := 0
	for root.Left != nil {
		d++
		root = root.Left
	}
	return d
}

func existsNode(pivot, d int, root *TreeNode) bool {
	for d > 0 {
		c := 1 << uint(d)
		if pivot < c/2 {
			root = root.Left
		} else {
			pivot -= c / 2
			root = root.Right
		}
		d--
	}
	return root != nil
}
