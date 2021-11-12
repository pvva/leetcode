package main

func main() {
	println(6, maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	println(4, maxProfit([]int{1, 2, 3, 4, 5}))
	println(0, maxProfit([]int{7, 6, 4, 3, 1}))
	println(0, maxProfit([]int{1}))
	println(1, maxProfit([]int{1, 2}))
}

/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/

You are given an array prices where prices[i] is the price of a given stock on the ith day.

Find the maximum profit you can achieve. You may complete at most two transactions.

Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

Example 1:

Input: prices = [3,3,5,0,0,3,1,4]
Output: 6
Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.

Example 2:

Input: prices = [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are engaging multiple transactions at
the same time. You must sell before buying again.

Example 3:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
Example 4:

Input: prices = [1]
Output: 0

Constraints:

1 <= prices.length <= 10^5
0 <= prices[i] <= 10^5
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	first_buy := -prices[0]
	first_sell := 0
	second_buy := -prices[0]
	second_sell := 0

	for _, price := range prices {
		first_sell = max(first_buy+price, first_sell)
		first_buy = max(-price, first_buy)
		second_sell = max(second_buy+price, second_sell)
		// mind that here we need to consider current first_sell max profit
		second_buy = max(first_sell-price, second_buy)
	}

	return max(first_sell, second_sell)
}
