package main

func main() {}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	totalProfit := 0

	for i := 1; i < len(prices); i++ {
		// If the current price is higher than the previous day, it's profitable to sell
		if prices[i] > prices[i-1] {
			totalProfit += prices[i] - prices[i-1]
		}
	}

	return totalProfit
}
