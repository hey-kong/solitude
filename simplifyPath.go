package main

import "strings"

// Leetcode 71. (medium)
func simplifyPath(path string) string {
	// path = "/a/../../b/../c//.//"
	nodes := strings.Split(path, "/")
	i := 0
	for i < len(nodes) {
		if nodes[i] == "" {
			nodes = append(nodes[:i], nodes[i+1:]...)
		} else {
			i++
		}
	}
	// nodes = [a .. .. b .. c .]
	stack := []string{}
	for _, node := range nodes {
		if node == ".." && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else if node != ".." && node != "." {
			stack = append(stack, node)
		}
	}
	// stack = [c]
	if len(stack) == 0 {
		return "/"
	}
	res := ""
	for _, node := range stack {
		res += "/" + node
	}
	return res
}
