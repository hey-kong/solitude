package main

// Leetcode 169. (easy)
func majorityElement(nums []int) int {
	res := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if cnt == 0 {
			res = nums[i]
			cnt = 1
		} else {
			if nums[i] == res {
				cnt++
			} else {
				cnt--
			}
		}
	}
	return res
}

// Leetcode 229. (medium)
func majorityElement2(nums []int) []int {
	a, b := 1<<31-1, 1<<31-1
	cnta, cntb := 0, 0
	for _, num := range nums {
		if num == a {
			cnta++
		} else if num == b {
			cntb++
		} else if cnta == 0 {
			a = num
			cnta = 1
		} else if cntb == 0 {
			b = num
			cntb = 1
		} else {
			cnta--
			cntb--
		}
	}
	cnta, cntb = 0, 0
	for _, num := range nums {
		if num == a {
			cnta++
		} else if num == b {
			cntb++
		}
	}
	res := []int{}
	if cnta > len(nums)/3 {
		res = append(res, a)
	}
	if cntb > len(nums)/3 {
		res = append(res, b)
	}
	return res
}
