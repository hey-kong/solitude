package main

import "strings"

// Leetcode 71. (medium)
func simplifyPath(path string) string {
	nodes := strings.Split(path, "/")
	stack := []string{}
	for _, node := range nodes {
		if node != "" && node != ".." && node != "." {
			stack = append(stack, node)
		} else if node == ".." && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return "/"
	}
	res := ""
	for _, node := range stack {
		res += "/" + node
	}
	return res
}
