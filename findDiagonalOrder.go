package main

// Leetcode 5394. (medium)
func findDiagonalOrder(nums [][]int) []int {
	cnt := 0
	for i := range nums {
		cnt += len(nums[i])
	}
	res := make([]int, cnt)
	i := 0
	queue := [][3]int{[3]int{0, 0, nums[0][0]}}
	for len(queue) != 0 {
		item := queue[0]
		queue = queue[1:]
		res[i] = item[2]
		i++

		// down
		if item[0]+1 < len(nums) && item[1] < len(nums[item[0]+1]) && nums[item[0]+1][item[1]] > 0 {
			queue = append(queue, [3]int{item[0] + 1, item[1], nums[item[0]+1][item[1]]})
			nums[item[0]+1][item[1]] = -nums[item[0]+1][item[1]]
		}
		// right
		if item[1]+1 < len(nums[item[0]]) {
			queue = append(queue, [3]int{item[0], item[1] + 1, nums[item[0]][item[1]+1]})
			nums[item[0]][item[1]+1] = -nums[item[0]][item[1]+1]
		}
	}
	return res
}
