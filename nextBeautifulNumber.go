package main

// Leetcode 5907. (medium)
func nextBeautifulNumber(n int) int {
	for {
		n++
		cur := n
		cnt := make([]int, 10)
		ok := true
		for cur != 0 {
			if cur%10 == 0 {
				ok = false
				break
			}
			cnt[cur%10]++
			cur /= 10
		}
		if ok {
			for i, c := range cnt {
				if c != 0 && i != c {
					ok = false
					break
				}
			}
		}
		if ok {
			return n
		}
	}
	return -1
}
