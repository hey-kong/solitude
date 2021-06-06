package main

// Leetcode 5776. (easy)
func findRotation(mat [][]int, target [][]int) bool {
	for i := 0; i < 4; i++ {
		if matIsEqual(mat, target) {
			return true
		}
		rotateImage(mat)
	}
	return false
}

func matIsEqual(mat1 [][]int, mat2 [][]int) bool {
	for i := range mat1 {
		for j := range mat2 {
			if mat1[i][j] != mat2[i][j] {
				return false
			}
		}
	}
	return true
}
