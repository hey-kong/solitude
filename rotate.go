package main

// Leetcode 189. (easy)
func rotate(nums []int, k int)  {
	if k >= len(nums) {
		k %= len(nums);
	}
	if k == 0 {
		return;
	}

	i := 0;
	start := 0;
	tmp := nums[i];
	for _ = range nums {
		i += k;
		if i >= len(nums) {
			i -= len(nums);
		}
		tmp, nums[i] = nums[i], tmp;

		if i == start {
			i++;
			start = i;
			tmp = nums[i]
		}
	}
}
