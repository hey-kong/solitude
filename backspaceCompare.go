package main

// Leetcode 844. (easy)
func backspaceCompare(S string, T string) bool {
	return buildOfBackspaceCompare(S) == buildOfBackspaceCompare(T)
}

func buildOfBackspaceCompare(str string) string {
	s := []byte{}
	for i := range str {
		if str[i] != '#' {
			s = append(s, str[i])
		} else if len(s) > 0 {
			s = s[:len(s)-1]
		}
	}
	return string(s)
}
