package main

// Leetcode 95. (medium)
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return dfsGenerateTrees(1, n)
}

func dfsGenerateTrees(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	res := []*TreeNode{}
	for i := start; i <= end; i++ {
		left := dfsGenerateTrees(start, i-1)
		right := dfsGenerateTrees(i+1, end)
		for _, l := range left {
			for _, r := range right {
				node := &TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				}
				res = append(res, node)
			}
		}
	}
	return res
}
