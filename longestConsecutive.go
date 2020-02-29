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
