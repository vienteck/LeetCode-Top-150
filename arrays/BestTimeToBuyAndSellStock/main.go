package main

func main() {
	maxProfit([]int{7, 6, 4, 3, 1})
}

func maxProfit(prices []int) int {

	buyingPoint := 0
	sellingPoint := 1
	maxProfit := 0
	for sellingPoint < len(prices) {

		if prices[sellingPoint]-prices[buyingPoint] > maxProfit {
			maxProfit = prices[sellingPoint] - prices[buyingPoint]
		} else if prices[sellingPoint] < prices[buyingPoint] {
			buyingPoint = sellingPoint
			sellingPoint++
		} else {
			sellingPoint++
		}
	}
	return maxProfit
}
