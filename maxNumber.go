package main

// Leetcode 321. (hard)
func maxNumber(nums1, nums2 []int, k int) (res []int) {
	start := 0
	if k > len(nums2) {
		start = k - len(nums2)
	}
	for i := start; i <= k && i <= len(nums1); i++ {
		s1 := subMaxNumber(nums1, i)
		s2 := subMaxNumber(nums2, k-i)
		merged := mergeMaxNumber(s1, s2)
		if lexicographicalLess(res, merged) {
			res = merged
		}
	}
	return
}

func subMaxNumber(nums []int, k int) []int {
	res := []int{}
	pop := len(nums) - k
	for i := 0; i < len(nums); i++ {
		for len(res) > 0 && res[len(res)-1] < nums[i] && pop > 0 {
			res = res[:len(res)-1]
			pop--
		}
		res = append(res, nums[i])
	}
	return res[:k]
}

func mergeMaxNumber(a, b []int) []int {
	merged := make([]int, len(a)+len(b))
	for k := range merged {
		if lexicographicalLess(a, b) {
			merged[k] = b[0]
			b = b[1:]
		} else {
			merged[k] = a[0]
			a = a[1:]
		}
	}
	return merged
}

func lexicographicalLess(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}
