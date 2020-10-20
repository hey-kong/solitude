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
