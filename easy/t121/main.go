package main

func maxProfit(prices []int) int {
	maxProfit := 0
	l := 0
	for i := 1; i < len(prices); i++ {
		maxProfit = max(maxProfit, prices[i]-prices[l])
		if prices[i] < prices[l] {
			l = i
		}
	}
	return maxProfit
}
