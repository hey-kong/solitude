package main

import "math"

// Leetcode 842. (medium)
func splitIntoFibonacci(S string) []int {
	n := len(S)
	res := []int{}
	var backtrack func(idx, pre, sum int) bool
	backtrack = func(idx, pre, sum int) bool {
		if idx == n {
			return len(res) >= 3
		}

		cur := 0
		for i := idx; i < n; i++ {
			if i > idx && S[idx] == '0' {
				break
			}

			cur = cur*10 + int(S[i]-'0')
			if cur > math.MaxInt32 {
				break
			}

			if len(res) >= 2 {
				if cur < sum {
					continue
				} else if cur > sum {
					break
				}
			}
			res = append(res, cur)
			if backtrack(i+1, cur, pre+cur) {
				return true
			}
			res = res[:len(res)-1]
		}
		return false
	}
	backtrack(0, 0, 0)
	return res
}
