package main

// Leetcode 4. (hard)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	left, right := (l1 + l2 + 1) / 2, (l1 + l2 + 2) / 2
	return float64(getKth(nums1, 0, l1 - 1, nums2, 0, l2 - 1, left) + getKth(nums1, 0, l1 - 1, nums2, 0, l2 - 1, right)) / 2;
}

func getKth(nums1 []int, start1 int, end1 int, nums2 []int, start2 int, end2 int, k int) int {
	l1 := end1 - start1 + 1
	l2 := end2 - start2 + 1
	if l1 > l2 {
		return getKth(nums2, start2, end2, nums1, start1, end1, k)
	}
	if l1 == 0 {
		return nums2[start2-1+k]
	}
	if k == 1 {
		return min(nums1[start1], nums2[start2])
	}

	i := start1 + min(l1, k/2) - 1
	j := start2 + min(l2, k/2) - 1
	if nums1[i] <= nums2[j] {
		return getKth(nums1, i+1, end1, nums2, start2, end2, k-(i-start1+1))
	} else {
		return getKth(nums1, start1, end1, nums2, j+1, end2, k-(j-start2+1))
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
