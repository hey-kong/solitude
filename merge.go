package main

// Leetcode 88. (easy)
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	mark := m + n - 1
	for mark >= 0 && i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[mark] = nums2[j]
			j--
		} else {
			nums1[mark] = nums1[i]
			i--
		}
		mark--
	}

	for mark >= 0 && j >= 0 {
		nums1[mark] = nums2[j]
		j--
		mark--
	}
}
