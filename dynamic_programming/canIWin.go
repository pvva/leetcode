package main

func main() {
	println(canIWin(10, 11), false)
	println(canIWin(4, 6), true)
	println(canIWin(10, 40), false)
	println(canIWin(10, 54), false)
}

/*
https://leetcode.com/problems/can-i-win/

In the "100 game," two players take turns adding, to a running total, any integer from 1..10. The player who
first causes the running total to reach or exceed 100 wins.

What if we change the game so that players cannot re-use integers?

For example, two players might take turns drawing from a common pool of numbers of 1..15 without replacement
until they reach a total >= 100.

Given an integer maxChoosableInteger and another integer desiredTotal, determine if the first player to move
can force a win, assuming both players play optimally.

You can always assume that maxChoosableInteger will not be larger than 20 and desiredTotal will not be larger than 300.

Example

Input:
maxChoosableInteger = 10
desiredTotal = 11

Output:
false

Explanation:
No matter which integer the first player choose, the first player will lose.
The first player can choose an integer from 1 up to 10.
If the first player choose 1, the second player can only choose integers from 2 up to 10.
The second player will win by choosing 10 and get a total = 11, which is >= desiredTotal.
Same with other integers chosen by the first player, the second player will always win.
*/

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if desiredTotal <= maxChoosableInteger {
		return true
	}
	if (maxChoosableInteger+1)*maxChoosableInteger/2 < desiredTotal {
		return false
	}

	return exploreWin(1<<maxChoosableInteger-1, maxChoosableInteger, desiredTotal, make(map[int]bool))
}

func exploreWin(mask, maxChoosableInteger, desiredTotal int, results map[int]bool) bool {
	if r, ex := results[mask]; ex {
		return r
	}
	v := maxChoosableInteger
	for i := 1; i <= mask; i <<= 1 {
		if mask&i == i && (v >= desiredTotal || !exploreWin(mask^i, maxChoosableInteger, desiredTotal-v, results)) {
			results[mask] = true
			return true
		}
		v--
	}

	results[mask] = false
	return false
}
