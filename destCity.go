package main

// Leetcode 5400. (easy)
func destCity(paths [][]string) string {
	m := make(map[string]string)
	for _, path := range paths {
		m[path[0]] = path[1]
	}
	for _, path := range paths {
		if _, ok := m[path[1]]; !ok {
			return path[1]
		}
	}
	return ""
}
