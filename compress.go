package main

import "strconv"

// Leetcode 309. (medium)
func compress(chars []byte) int {
	mark := 0
	num := 0
	for i := range chars {
		num++
		if i < len(chars)-1 && chars[i] == chars[i+1] {
			continue
		}

		if num == 1 {
			chars[mark] = chars[i]
			mark++
			num = 0
		} else {
			chars[mark] = chars[i]
			mark++
			tmp := []byte(strconv.Itoa(num))
			for j := range tmp {
				chars[mark] = tmp[j]
				mark++
			}
			num = 0
		}
	}
	return mark
}
