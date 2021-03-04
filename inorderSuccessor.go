package main

// Leetcode 285. (medium)
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	return recursiveInorderSuccessor(root, p, nil)
}

func recursiveInorderSuccessor(root, p, res *TreeNode) *TreeNode {
	if root == nil || res != nil {
		return res
	}
	res = recursiveInorderSuccessor(root.Left, p, res)
	if res == nil && root.Val > p.Val {
		res = root
	}
	res = recursiveInorderSuccessor(root.Right, p, res)
	return res
}

// Leetcode 510. (medium)
func inorderSuccessor2(node *TreeNode2) *TreeNode2 {
	if node.Right != nil {
		node = node.Right
		for node.Left != nil {
			node = node.Left
		}
		return node
	}

	for node.Parent != nil && node.Parent.Right == node {
		node = node.Parent
	}
	return node.Parent
}
