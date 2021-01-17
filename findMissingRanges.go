package main

import (
	"fmt"
	"strconv"
)

// Leetcode 163. (medium)
func findMissingRanges(nums []int, lower int, upper int) (res []string) {
	if len(nums) == 0 {
		if lower != upper {
			return []string{fmt.Sprintf("%s->%s", strconv.Itoa(lower), strconv.Itoa(upper))}
		}
		return []string{fmt.Sprintf("%s", strconv.Itoa(lower))}
	}

	if lower != nums[0] {
		if lower+1 == nums[0] {
			res = append(res, strconv.Itoa(lower))
		} else {
			res = append(res, fmt.Sprintf("%s->%s", strconv.Itoa(lower), strconv.Itoa(nums[0]-1)))
		}
	}
	pre := nums[0]
	for _, num := range nums {
		if num == pre || num == pre+1 {
			pre = num
			continue
		}

		if pre+2 == num {
			res = append(res, strconv.Itoa(pre+1))
		} else {
			res = append(res, fmt.Sprintf("%s->%s", strconv.Itoa(pre+1), strconv.Itoa(num-1)))
		}
		pre = num
	}
	if pre != upper {
		if pre+1 == upper {
			res = append(res, strconv.Itoa(upper))
		} else {
			res = append(res, fmt.Sprintf("%s->%s", strconv.Itoa(pre+1), strconv.Itoa(upper)))
		}
	}
	return
}
