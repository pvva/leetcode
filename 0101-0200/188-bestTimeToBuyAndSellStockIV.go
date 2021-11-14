package main

func main() {
	println(2, maxProfit(2, []int{2, 4, 1}))
	println(7, maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
	println(15, maxProfit(4, []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}))
}

/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/

You are given an integer array prices where prices[i] is the price of a given stock on the ith day, and an integer k.
Find the maximum profit you can achieve. You may complete at most k transactions.

Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

Example 1:

Input: k = 2, prices = [2,4,1]
Output: 2
Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit = 4-2 = 2.

Example 2:

Input: k = 2, prices = [3,2,6,5,0,3]
Output: 7
Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit = 6-2 = 4. Then buy on day 5 (price = 0)
and sell on day 6 (price = 3), profit = 3-0 = 3.

Constraints:

0 <= k <= 100
0 <= prices.length <= 1000
0 <= prices[i] <= 1000
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	buys := make([]int, k)
	sells := make([]int, k)
	currentMax := 0

	for i := 0; i < k; i++ {
		buys[i] = -prices[0]
	}

	for _, price := range prices {
		currentProfit := 0
		for j := 0; j < k; j++ {
			sells[j] = max(buys[j]+price, sells[j])
			buys[j] = max(currentProfit-price, buys[j])

			currentProfit = sells[j]
			if currentMax < sells[j] {
				currentMax = sells[j]
			}
		}
	}

	return currentMax
}
