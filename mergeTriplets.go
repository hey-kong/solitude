package main

// Leetcode 5785. (medium)
func mergeTriplets(triplets [][]int, target []int) bool {
	t := make([]int, 3)
	for _, arr := range triplets {
		if arr[0] <= target[0] && arr[1] <= target[1] && arr[2] <= target[2] {
			t[0] = max(t[0], arr[0])
			t[1] = max(t[1], arr[1])
			t[2] = max(t[2], arr[2])
		}
	}
	return t[0] == target[0] && t[1] == target[1] && t[2] == target[2]
}
