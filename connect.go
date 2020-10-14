package main

type Node1 struct {
	Val   int
	Left  *Node1
	Right *Node1
	Next  *Node1
}

// Leetcode 116. (medium)
func connect(root *Node1) *Node1 {
	res := recursiveConnect(root, 0, [][]*Node1{})
	for i := range res {
		for j := range res[i] {
			if j+1 != len(res[i]) {
				res[i][j].Next = res[i][j+1]
			}
		}
	}
	return root
}

func recursiveConnect(root *Node1, depth int, res [][]*Node1) [][]*Node1 {
	if root == nil {
		return res
	}
	if len(res) == depth {
		res = append(res, []*Node1{})
	}
	res[depth] = append(res[depth], root)
	res = recursiveConnect(root.Left, depth+1, res)
	res = recursiveConnect(root.Right, depth+1, res)
	return res
}
