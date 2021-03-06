package main

// Leetcode 121. (easy)
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	minprice := prices[0]
	maxprofit := 0
	for i := 1; i < len(prices); i++ {
		maxprofit = max(maxprofit, prices[i]-minprice)
		minprice = min(minprice, prices[i])
	}
	return maxprofit
}

// Leetcode 122. (easy)
func maxProfit2(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

// Leetcode 123. (hard)
func maxProfit3(prices []int) int {
	return maxProfit4(2, prices)
}

// Leetcode 188. (hard)
func maxProfit4(k int, prices []int) int {
	if len(prices) == 0 || k == 0 {
		return 0
	}

	k = min(k, len(prices)/2)
	hold, sell := make([][]int, k+1), make([][]int, k+1)
	for i := 0; i <= k; i++ {
		hold[i] = make([]int, len(prices))
		sell[i] = make([]int, len(prices))
	}
	hold[0][0] = -prices[0]
	for i := 1; i <= k; i++ {
		hold[i][0], sell[i][0] = -1000000, -1000000
	}

	for j := 1; j < len(prices); j++ {
		hold[0][j], sell[0][j] = max(hold[0][j-1], -prices[j]), 0
		for i := 1; i <= k; i++ {
			hold[i][j] = max(hold[i][j-1], sell[i][j-1]-prices[j])
			sell[i][j] = max(sell[i][j-1], hold[i-1][j-1]+prices[j])
		}
	}
	res := sell[0][len(prices)-1]
	for i := 1; i <= k; i++ {
		res = max(res, sell[i][len(prices)-1])
	}
	return res
}

// Leetcode 309. (medium)
func maxProfitWithCooldown(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][2])
		dp[i][2] = dp[i-1][0] + prices[i]
	}
	return max(dp[len(prices)-1][1], dp[len(prices)-1][2])
}

// Leetcode 714. (medium)
func maxProfitWithFee(prices []int, fee int) int {
	hold, sell := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		hold, sell = max(hold, sell-prices[i]), max(sell, hold+prices[i]-fee)
	}
	return max(hold, sell)
}
