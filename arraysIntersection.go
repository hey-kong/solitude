package main

// Leetcode 1213. (easy)
func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	res := []int{}
	i1, i2, i3 := 0, 0, 0
	n1, n2, n3 := len(arr1), len(arr2), len(arr3)
	for i1 < n1 && i2 < n2 && i3 < n3 {
		if arr1[i1] == arr2[i2] && arr2[i2] == arr3[i3] {
			res = append(res, arr1[i1])
			i1++
			i2++
			i3++
		} else {
			m := min(arr3[i3], min(arr1[i1], arr2[i2]))
			if arr1[i1] == m {
				i1++
			}
			if arr2[i2] == m {
				i2++
			}
			if arr3[i3] == m {
				i3++
			}
		}
	}
	return res
}
