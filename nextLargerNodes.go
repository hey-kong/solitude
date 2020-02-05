package main

// Leetcode 1019. (medium)
func nextLargerNodes(head *ListNode) []int {
	nums := listToArray(head)
	res := make([]int, len(nums))
	stack := []int{}
	for i, num := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < num {
			res[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}

func listToArray(head *ListNode) []int {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}
