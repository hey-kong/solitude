package main

import "strconv"

// Leetcode 165. (medium)
func compareVersion(version1 string, version2 string) int {
	i, j := 0, 0
	for i < len(version1) || j < len(version2) {
		start1 := i
		for i < len(version1) && version1[i] != '.' {
			i++
		}
		for start1 < i && version1[start1] == '0' {
			start1++
		}

		start2 := j
		for j < len(version2) && version2[j] != '.' {
			j++
		}
		for start2 < j && version2[start2] == '0' {
			start2++
		}

		num1, num2 := 0, 0
		if start1 < i {
			num1, _ = strconv.Atoi(version1[start1:i])
		}
		if start2 < j {
			num2, _ = strconv.Atoi(version2[start2:j])
		}

		if num1 < num2 {
			return -1
		} else if num1 > num2 {
			return 1
		}

		i++
		j++
	}
	return 0
}
