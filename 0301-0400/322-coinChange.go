package main

func main() {
	println(coinChange([]int{1, 2, 5}, 11), 3)
	println(coinChange([]int{2, 5, 10, 1}, 27), 4)
	println(coinChange([]int{1, 2, 5}, 100), 20)
}

/*
https://leetcode.com/problems/coin-change/

You are given coins of different denominations and a total amount of money amount. Write a function to compute the
fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination
of the coins, return -1.

Example 1:

Input: coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

Example 2:

Input: coins = [2], amount = 3
Output: -1
Note:
You may assume that you have an infinite number of each kind of coin.
*/

func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}

	bestSteps := make([]int, amount+1)
	for i := range bestSteps {
		bestSteps[i] = -1
	}
	bestSteps[0] = 0

	for a := 1; a <= amount; a++ {
		for _, c := range coins {
			if c > a {
				continue
			}
			t := bestSteps[a-c] + 1
			if t > 0 && (bestSteps[a] == -1 || t < bestSteps[a]) {
				bestSteps[a] = t
			}
		}
	}

	return bestSteps[amount]
}
