package main

// Leetcode 152. (medium)
func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp_max := make([]int, len(nums))
	dp_min := make([]int, len(nums))
	dp_max[0], dp_min[0] = nums[0], nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		preMax, preMin := dp_max[i-1], dp_min[i-1]
		if nums[i] < 0 {
			preMax, preMin = preMin, preMax
		}
		dp_max[i] = max(nums[i], preMax*nums[i])
		dp_min[i] = min(nums[i], preMin*nums[i])
		res = max(res, dp_max[i])
	}
	return res
}

// Leetcode 1339. (medium)
func maxProduct2(root *TreeNode) int {
	sum := calPathSum(root, 0)
	_, res := recursiveMaxProduct2(root, sum, 0)
	return res % (1e9 + 7)
}

func recursiveMaxProduct2(root *TreeNode, sum, res int) (int, int) {
	if root == nil {
		return 0, res
	}
	leftSum, leftMax := recursiveMaxProduct2(root.Left, sum, res)
	rightSum, rightMax := recursiveMaxProduct2(root.Right, sum, res)
	p1 := leftSum * (sum - leftSum)
	p2 := rightSum * (sum - rightSum)
	res = max(max(leftMax, rightMax), max(p1, p2))
	return leftSum + rightSum + root.Val, res
}

func calPathSum(root *TreeNode, sum int) int {
	if root == nil {
		return sum
	}
	sum += root.Val
	sum = calPathSum(root.Left, sum)
	sum = calPathSum(root.Right, sum)
	return sum
}
