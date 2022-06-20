package main

// Leetcode 2312. (hard)
func sellingWood(m int, n int, prices [][]int) int64 {
	woodPrices := make([][]int, m+1)
	for i := range woodPrices {
		woodPrices[i] = make([]int, n+1)
	}
	for _, price := range prices {
		woodPrices[price[0]][price[1]] = price[2]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k <= j/2; k++ {
				woodPrices[i][j] = max(woodPrices[i][j], woodPrices[i][k]+woodPrices[i][j-k])
			}
			for k := 1; k <= i/2; k++ {
				woodPrices[i][j] = max(woodPrices[i][j], woodPrices[k][j]+woodPrices[i-k][j])
			}
		}
	}
	return int64(woodPrices[m][n])
}
