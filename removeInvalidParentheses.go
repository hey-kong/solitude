package main

// Leetcode 301. (hard)
func removeInvalidParentheses(s string) []string {
	l, r := 0, 0
	for i := range s {
		if s[i] == '(' {
			l++
		} else if s[i] == ')' {
			if l == 0 {
				r++
			} else {
				l--
			}
		}
	}

	return dfsRemoveInvalidParentheses(l, r, 0, 0, s, "", []string{})
}

func dfsRemoveInvalidParentheses(l, r, cnt, i int, s, tmp string, res []string) []string {
	if i == len(s) {
		if l == 0 && r == 0 && cnt == 0 {
			res = append(res, tmp)
		}
		return res
	}

	if s[i] == '(' {
		j := i
		for j < len(s) && s[j] == '(' {
			j++
		}
		l -= j - i
		for k := i; k <= j; k++ {
			if l >= 0 {
				res = dfsRemoveInvalidParentheses(l, r, cnt, j, s, tmp, res)
			}
			l++
			cnt++
			tmp += "("
		}
	} else if s[i] == ')' {
		j := i
		for j < len(s) && s[j] == ')' {
			j++
		}
		r -= j - i
		for k := i; k <= j; k++ {
			if cnt >= 0 && r >= 0 {
				res = dfsRemoveInvalidParentheses(l, r, cnt, j, s, tmp, res)
			}
			r++
			cnt--
			tmp += ")"
		}
	} else {
		res = dfsRemoveInvalidParentheses(l, r, cnt, i+1, s, tmp+s[i:i+1], res)
	}
	return res
}
