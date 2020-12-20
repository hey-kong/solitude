package main

// Leetcode 116. (medium)
func connect(root *TreeNode3) *TreeNode3 {
	res := recursiveConnect(root, 0, [][]*TreeNode3{})
	for i := range res {
		for j := range res[i] {
			if j+1 != len(res[i]) {
				res[i][j].Next = res[i][j+1]
			}
		}
	}
	return root
}

func recursiveConnect(root *TreeNode3, depth int, res [][]*TreeNode3) [][]*TreeNode3 {
	if root == nil {
		return res
	}
	if len(res) == depth {
		res = append(res, []*TreeNode3{})
	}
	res[depth] = append(res[depth], root)
	res = recursiveConnect(root.Left, depth+1, res)
	res = recursiveConnect(root.Right, depth+1, res)
	return res
}

// Leetcode 117. (medium)
func connect2(root *TreeNode3) *TreeNode3 {
	start := root
	for start != nil {
		var nextStart, last *TreeNode3
		dfs := func(cur *TreeNode3) {
			if cur == nil {
				return
			}
			if nextStart == nil {
				nextStart = cur
			}
			if last != nil {
				last.Next = cur
			}
			last = cur
		}
		for p := start; p != nil; p = p.Next {
			dfs(p.Left)
			dfs(p.Right)
		}
		start = nextStart
	}
	return root
}
