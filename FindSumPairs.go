package main

// Leetcode 5761. (medium)
type FindSumPairs struct {
	nums1 []int
	nums2 []int
	m     map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	m := make(map[int]int, len(nums2))
	for _, num := range nums2 {
		m[num]++
	}

	return FindSumPairs{
		nums1,
		nums2,
		m,
	}
}

func (this *FindSumPairs) Add(index int, val int) {
	this.m[this.nums2[index]]--
	this.nums2[index] += val
	this.m[this.nums2[index]]++
}

func (this *FindSumPairs) Count(tot int) int {
	cnt := 0
	for _, num := range this.nums1 {
		cnt += this.m[tot-num]
	}
	return cnt
}
