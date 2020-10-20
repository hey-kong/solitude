package main

// Leetcode 349. (easy)
func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	res := []int{}
	for _, num := range nums1 {
		m[num] = true
	}
	for _, num := range nums2 {
		if ok := m[num]; ok {
			res = append(res, num)
			m[num] = false
		}
	}
	return res
}
