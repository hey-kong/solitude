package main

// Leetcode 2306. (hard)
func distinctNames(ideas []string) int64 {
	group := make([]map[string]bool, 26)
	for i := range group {
		group[i] = make(map[string]bool)
	}

	for _, idea := range ideas {
		group[idea[0]-'a'][idea[1:]] = true
	}

	res := int64(0)
	for i, a := range group {
		for _, b := range group[i+1:] {
			d := 0
			for s := range a {
				if b[s] {
					d++
				}
			}
			res += int64(len(a)-d) * int64(len(b)-d)
		}
	}
	return res * 2
}
