package main

// Leetcode 1071. (easy)
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	l := gcd(len(str1), len(str2))
	return str1[:l]
}
