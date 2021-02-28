package main

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */

// Leetcode 277. (medium)
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		celebrity := 0
		for i := 1; i < n; i++ {
			// if a knows b, a is not the celebrity
			// if a doesn't know b, b is not the celebrity
			if knows(celebrity, i) {
				celebrity = i
			}
		}
		for i := 0; i < n; i++ {
			if i != celebrity && (knows(celebrity, i) || !knows(i, celebrity)) {
				return -1
			}
		}
		return celebrity
	}
}
