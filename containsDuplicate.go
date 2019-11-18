package main

// Leetcode 217. (easy)
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = true
	}
	return false
}

// Leetcode 219. (easy)
func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 {
		return false
	}

	m := make(map[int]bool, k)
	for i, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		if len(m) == k {
			delete(m, nums[i-k])
		}
		m[num] = true
	}
	return false
}

// Leetcode 220. (medium)
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k == 0 || t < 0 {
		return false
	}

	allBucket := map[int]int{}
	bucketSize := t + 1
	for i, num := range nums {
		bucketId := getID(num, bucketSize)
		if _, ok := allBucket[bucketId]; ok {
			return true
		}
		if v, ok := allBucket[bucketId-1]; ok && num-v <= t {
			return true
		}
		if v, ok := allBucket[bucketId+1]; ok && v-num <= t {
			return true
		}
		allBucket[bucketId] = num
		if i >= k {
			delete(allBucket, getID(nums[i-k], bucketSize))
		}
	}
	return false
}

func getID(x, w int) int {
	if x < 0 {
		return x/w - 1
	}
	return x/w
}
