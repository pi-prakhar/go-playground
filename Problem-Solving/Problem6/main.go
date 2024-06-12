package main

/**
PROBLEM:
Leetcode -> Best Time to Buy and Sell Stock -> easy

TAGS:
arrays, two pointer, sliding window

**/

func maxProfit(prices []int) int {
	buy := prices[0]
	profit := 0
	for _, price := range prices {
		if price < buy {
			buy = price
			continue
		}
		if (price - buy) > profit {
			profit = price - buy
		}
	}
	return profit
}
