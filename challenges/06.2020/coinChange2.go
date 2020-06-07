package main

func main() {
	println(change(5, []int{1, 2, 5}), 4)
	println(change(3, []int{2}), 0)
	println(change(10, []int{10}), 1)
}

/*
You are given coins of different denominations and a total amount of money. Write a function to compute
the number of combinations that make up that amount. You may assume that you have infinite number of each kind of coin.

Example 1:

Input: amount = 5, coins = [1, 2, 5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1

Example 2:

Input: amount = 3, coins = [2]
Output: 0
Explanation: the amount of 3 cannot be made up just with coins of 2.

Example 3:

Input: amount = 10, coins = [10]
Output: 1
*/

func change(amount int, coins []int) int {
	memo := make([][]int, len(coins)+1)
	for x := range memo {
		memo[x] = make([]int, amount+1)
	}

	memo[0][0] = 1

	for x, coin := range coins {
		memo[x+1][0] = 1
		for y := 1; y <= amount; y++ {
			memo[x+1][y] = memo[x][y] // exclude current coin - take the amount of previous coin
			if coin <= y {
				memo[x+1][y] += memo[x+1][y-coin] // try to include current coin
			}
		}
	}

	return memo[len(coins)][amount]
}
