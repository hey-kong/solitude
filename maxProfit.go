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
func maxProfit3(prices []int, fee int) int {
	hold, sell := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		hold, sell = max(hold, sell-prices[i]), max(sell, hold+prices[i]-fee)
	}
	return max(hold, sell)
}
