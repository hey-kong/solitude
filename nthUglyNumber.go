package main

// Leetcode 264. (medium)
func nthUglyNumber(n int) int {
	nums := make([]int, n)
	nums[0] = 1
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < n; i++ {
		num := min(min(nums[i2]*2, nums[i3]*3), nums[i5]*5)
		nums[i] = num
		if nums[i2]*2 == num {
			i2++
		}
		if nums[i3]*3 == num {
			i3++
		}
		if nums[i5]*5 == num {
			i5++
		}
	}
	return nums[n-1]
}

// Leetcode 1201. (medium)
func nthUglyNumber2(n int, a int, b int, c int) int {
	ab, ac, bc := lcm(a, b), lcm(a, c), lcm(b, c)
	abc := lcm(ab, c)
	left, right := 0, 1<<31-1
	for left <= right {
		mid := left + (right-left)/2
		cnt := mid/a + mid/b + mid/c - mid/ab - mid/ac - mid/bc + mid/abc
		if cnt == n {
			right = mid - 1
		} else if cnt < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func gcd(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		tmp := b
		b = a % b
		a = tmp
	}
	return a
}

func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return a * b / gcd(a, b)
}
