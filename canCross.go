package main

// Leetcode 403. (hard)
func canCross(stones []int) bool {
	m := make(map[int][]int)
	for _, s := range stones {
		m[s] = []int{}
	}
	m[0] = append(m[0], 0)

	for _, s := range stones {
		isRepeated := make(map[int]bool)
		for _, step := range m[s] {
			if isRepeated[step] {
				continue
			} else {
				isRepeated[step] = true
			}

			for i := step + 1; i > 0 && i >= step-1; i-- {
				if _, ok := m[s+i]; ok {
					m[s+i] = append(m[s+i], i)
				}
			}
		}
		if len(m[stones[len(stones)-1]]) != 0 {
			return true
		}
	}
	return false
}
