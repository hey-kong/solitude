package main

// Leetcode 307. (medium)
type NumArray struct {
	arr  []int
	tree []int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{arr: nums, tree: nums}
	}
	tree := make([]int, 4*len(nums))
	buildTree(nums, tree, 1, 0, len(nums)-1)
	return NumArray{arr: nums, tree: tree}
}

func (this *NumArray) Update(i int, val int) {
	updateTree(this.arr, this.tree, 1, 0, len(this.arr)-1, i, val)
}

func (this *NumArray) SumRange(i int, j int) int {
	if i > j {
		return 0
	}
	return queryTree(this.arr, this.tree, 1, 0, len(this.arr)-1, i, j)
}

func buildTree(arr []int, tree []int, node, start, end int) {
	if start == end {
		tree[node] = arr[start]
		return
	}

	leftNode := 2 * node
	rightNode := 2*node + 1
	mid := start + (end-start)/2
	buildTree(arr, tree, leftNode, start, mid)
	buildTree(arr, tree, rightNode, mid+1, end)
	tree[node] = tree[leftNode] + tree[rightNode]
}

func updateTree(arr []int, tree []int, node, start, end, index, value int) {
	if start == end {
		arr[index] = value
		tree[node] = value
		return
	}

	leftNode := 2 * node
	rightNode := 2*node + 1
	mid := start + (end-start)/2
	if index >= start && index <= mid {
		updateTree(arr, tree, leftNode, start, mid, index, value)
	} else if index >= mid+1 && index <= end {
		updateTree(arr, tree, rightNode, mid+1, end, index, value)
	}
	tree[node] = tree[leftNode] + tree[rightNode]
}

func queryTree(arr []int, tree []int, node, start, end, L, R int) int {
	if R < start || L > end {
		return 0
	} else if L <= start && R >= end {
		return tree[node]
	}

	leftNode := 2 * node
	rightNode := 2*node + 1
	mid := start + (end-start)/2
	sum1 := queryTree(arr, tree, leftNode, start, mid, L, R)
	sum2 := queryTree(arr, tree, rightNode, mid+1, end, L, R)
	return sum1 + sum2
}
