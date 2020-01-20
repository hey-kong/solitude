package main

// Leetcode 678. (medium)
func checkValidString(s string) bool {
    left := make([]int, len(s))
	star := make([]int, len(s))
	i, j := 0, 0
	for k, r := range s {
		if r == '(' {
			left[i] = k
			i++
		} else if r == '*' {
			star[j] = k
			j++
		} else if r == ')' {
			if i > 0 {
				i--
			} else if j > 0 {
				j--
			} else {
				return false
			}
		}
	}
	for i != 0 {
		if j == 0 {
			return false
		}
		i--
		j--
		if left[i] > star[j] {
			return false
		}
	}
	return true
}
