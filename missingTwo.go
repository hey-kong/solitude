package main

// Leetcode m17.19. (hard)
func missingTwo(nums []int) []int {
	n := len(nums) + 2
	res := 0
	for i := 1; i <= n; i++ {
		res ^= i
	}
	for _, num := range nums {
		res ^= num
	}

	lowbit := res & (-res)
	one := 0
	for i := 1; i <= n; i++ {
		if i&lowbit == 0 {
			one ^= i
		}
	}
	for _, num := range nums {
		if num&lowbit == 0 {
			one ^= num
		}
	}
	return []int{one, one ^ res}
}

func missingTwo2(nums []int) []int {
	n := len(nums) + 2
	sum := (1 + n) * n / 2
	for _, num := range nums {
		sum += num
	}

	sumTwo := (1+n)*n - sum
	average := sumTwo / 2
	sum = 0
	for _, num := range nums {
		if num <= average {
			sum += num
		}
	}
	one := (1+average)*average/2 - sum
	return []int{one, sumTwo - one}
}

func missingTwo3(nums []int) []int {
	nums = append(nums, []int{-1, -1, -1}...)
	for i := 0; i < len(nums); i++ {
		for i != nums[i] && nums[i] != -1 {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}

	res := []int{}
	for i := 1; i < len(nums); i++ {
		if nums[i] == -1 {
			res = append(res, i)
		}
	}
	return res
}
