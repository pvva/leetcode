package main

func main() {
	println(maxProfit([]int{1, 2, 3, 0, 2}))
	println(maxProfit([]int{1, 11, 2, 7, 4}))
	println(maxProfit([]int{6, 1, 6, 4, 3, 0, 2}))
	println(maxProfit([]int{1, 4, 7, 5, 6, 2, 5, 1, 9, 7, 9, 7, 0, 6, 8}))
	println(maxProfitLinear([]int{1, 4, 7, 5, 6, 2, 5, 1, 9, 7, 9, 7, 0, 6, 8}))
}

/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like
(ie, buy one and sell one share of the stock multiple times) with the following restrictions:

You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)

Example:

Input: [1,2,3,0,2]
Output: 3
Explanation: transactions = [buy, sell, cooldown, buy, sell]
*/

func maxProfit(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}

	profit := make([]int, l)
	profit[1] = max(0, prices[1]-prices[0])

	for i := 2; i < l; i++ {
		p := 0
		for j := i - 1; j >= 0 && prices[j] < prices[i]; j-- {
			s := 0
			if j > 1 {
				s = profit[j-2]
			}
			p = max(p, prices[i]-prices[j]+s)
		}

		profit[i] = max(profit[i-1], p)
	}

	return profit[l-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProfitLinear(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}

	// the most profits for the last action 'sell', till i's day
	sells := make([]int, l)
	// the most profits for the last action 'buy', till i's day
	buys := make([]int, l)

	buys[0] = -prices[0]

	sells[1] = max(0, prices[1]-prices[0])
	buys[1] = max(-prices[0], -prices[1])

	for i := 2; i < l; i++ {
		// max of buy at prev day + sell now _and_ sell at prev day
		sells[i] = max(buys[i-1]+prices[i], sells[i-1])
		// max of sell -2 days (due to 1 day of a cooldown) _and_ buy at prev day
		buys[i] = max(sells[i-2]-prices[i], buys[i-1])
	}

	return sells[l-1]
}
