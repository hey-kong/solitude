package main

// Leetcode 314. (medium)
func verticalOrder(root *TreeNode) [][]int {
	leftArr := []valPos{valPos{}}
	rightArr := []valPos{}
	leftArr, rightArr = recursiveVerticalOrder(root, 0, 0, leftArr, rightArr)
	res := [][]int{}
	for i := len(leftArr) - 1; i > 0; i-- {
		tmp := []int{}
		valPosBubbleSort(leftArr[i])
		for j := range leftArr[i] {
			tmp = append(tmp, leftArr[i][j][0])
		}
		res = append(res, tmp)
	}
	for i := 0; i < len(rightArr); i++ {
		tmp := []int{}
		valPosBubbleSort(rightArr[i])
		for j := range rightArr[i] {
			tmp = append(tmp, rightArr[i][j][0])
		}
		res = append(res, tmp)
	}
	return res
}

func recursiveVerticalOrder(root *TreeNode, i, j int, leftArr, rightArr []valPos) ([]valPos, []valPos) {
	if root == nil {
		return leftArr, rightArr
	}
	if i < 0 {
		if len(leftArr) == -i {
			leftArr = append(leftArr, valPos{})
		}
		leftArr[-i] = append(leftArr[-i], [2]int{root.Val, j})
	} else {
		if len(rightArr) == i {
			rightArr = append(rightArr, valPos{})
		}
		rightArr[i] = append(rightArr[i], [2]int{root.Val, j})
	}
	leftArr, rightArr = recursiveVerticalOrder(root.Left, i-1, j+1, leftArr, rightArr)
	leftArr, rightArr = recursiveVerticalOrder(root.Right, i+1, j+1, leftArr, rightArr)
	return leftArr, rightArr
}

type valPos [][2]int

func valPosBubbleSort(arr valPos) valPos {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 1; j <= i; j++ {
			if arr[j-1][1] > arr[j][1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return arr
}
