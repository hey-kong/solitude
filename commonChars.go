package main

// Leetcode 1002. (easy)
func commonChars(A []string) []string {
	arr := make([][]int, len(A))
	for i := range arr {
		arr[i] = make([]int, 26)
	}
	for i, s := range A {
		for _, b := range s {
			arr[i][b-'a']++
		}
	}
	res := []string{}
	n := len(A)
	for i := 0; i < 26; i++ {
		m := 100
		for j := 0; j < n; j++ {
			m = min(m, arr[j][i])
		}
		if m == 0 {
			continue
		}
		for k := 0; k < m; k++ {
			res = append(res, string('a'+i))
		}
	}
	return res
}
