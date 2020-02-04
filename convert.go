package main

// Leetcode 6. (medium)
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	strs := make([][]byte, numRows)
	for i := range strs {
		strs[i] = []byte{}
	}
	j := 0
	flag := 1
	for i := range s {
		strs[j] = append(strs[j], s[i])
		if flag == 1 && j == numRows-1 {
			flag = -1
		}
		if flag == -1 && j == 0 {
			flag = 1
		}
		j += flag
	}
	for i := 1; i < numRows; i++ {
		strs[0] = append(strs[0], strs[i]...)
	}
	return string(strs[0])
}
