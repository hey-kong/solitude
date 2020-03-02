package main

// Leetcode 1287. (easy)
func findSpecialInteger(arr []int) int {
	scan := len(arr) / 4
	for i := 0; i < len(arr); i += scan {
		left := leftBound(arr, arr[i])
		right := rightBound(arr, arr[i])
		if right-left+1 > len(arr)/4 {
			return arr[i]
		}
	}
	return -1
}
