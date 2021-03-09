package main

// Leetcode 128. (hard)
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	root, size := make(map[int]int, len(nums)), make(map[int]int, len(nums))
	for _, num := range nums {
		root[num] = num
		size[num] = 1
	}
	res := 1
	tmp := 0
	for _, num := range nums {
		if _, ok := root[num+1]; ok && num != 1<<31 {
			tmp, root, size = unionOf128(num, num+1, root, size)
			res = max(res, tmp)
		}
	}
	return res
}

func findOf128(x int, root map[int]int) int {
	for x != root[x] {
		x = root[x]
	}
	return x
}

func unionOf128(x, y int, root, size map[int]int) (int, map[int]int, map[int]int) {
	x = findOf128(x, root)
	y = findOf128(y, root)
	if x == y {
		return size[y], root, size
	}
	root[x] = y
	size[y] += size[x]
	return size[y], root, size
}

func longestConsecutive2(nums []int) int {
	m := make(map[int]int, len(nums))
	res := 0
	for _, num := range nums {
		if _, ok := m[num]; ok {
			continue
		}

		left, right := 0, 0
		if _, ok := m[num-1]; ok {
			left = m[num-1]
		}
		if _, ok := m[num+1]; ok {
			right = m[num+1]
		}
		cur := left + 1 + right
		res = max(res, cur)
		m[num-left], m[num], m[num+right] = cur, cur, cur
	}
	return res
}

// Leetcode 298. (medium)
func longestConsecutive3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return recursiveLongestConsecutive3(root, root.Val, 0)
}

func recursiveLongestConsecutive3(root *TreeNode, target, cur int) int {
	if root == nil {
		return cur
	}

	if root.Val == target {
		res1 := recursiveLongestConsecutive3(root.Left, root.Val+1, cur+1)
		res2 := recursiveLongestConsecutive3(root.Right, root.Val+1, cur+1)
		return max(res1, res2)
	}
	res1 := recursiveLongestConsecutive3(root.Left, root.Val+1, 1)
	res2 := recursiveLongestConsecutive3(root.Right, root.Val+1, 1)
	return max(cur, max(res1, res2))
}
