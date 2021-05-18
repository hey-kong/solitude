package main

// Leetcode 164. (hard)
func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	maxVal := maxInt(nums...)
	buf := make([][]int, 10)
	for exp := 1; exp <= maxVal; exp *= 10 {
		for _, num := range nums {
			bucket := num / exp % 10
			buf[bucket] = append(buf[bucket], num)
		}

		k := 0
		for i := range buf {
			for j := range buf[i] {
				nums[k] = buf[i][j]
				k++
			}
		}

		buf = make([][]int, 10)
	}

	res := 0
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] > res {
			res = nums[i] - nums[i-1]
		}
	}
	return res
}

func maxInt(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
