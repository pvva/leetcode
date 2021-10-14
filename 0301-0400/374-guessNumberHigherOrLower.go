package main

import "fmt"

var pick int

func main() {
	for _, tc := range []struct {
		n    int
		pick int
	}{
		{n: 10, pick: 6},
		{n: 1, pick: 1},
		{n: 2, pick: 1},
		{n: 2, pick: 2},
	} {
		pick = tc.pick
		fmt.Printf("n=%d, pick=%d, guess=%d\n", tc.n, tc.pick, guessNumber(tc.n))
	}
}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 */
func guess(num int) int {
	if num < pick {
		return -1
	}
	if num > pick {
		return 1
	}
	return 0
}

/*
https://leetcode.com/problems/guess-number-higher-or-lower/

We are playing the Guess Game. The game is as follows:

I pick a number from 1 to n. You have to guess which number I picked.

Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.

You call a pre-defined API int guess(int num), which returns 3 possible results:

-1: The number I picked is lower than your guess (i.e. pick < num).
1: The number I picked is higher than your guess (i.e. pick > num).
0: The number I picked is equal to your guess (i.e. pick == num).
Return the number that I picked.
*/

func guessNumber(n int) int {
	l := 1
	h := n
	for l <= h {
		m := l + (h-l)/2
		switch guess(m) {
		case 0:
			return m
		case -1:
			l = m + 1
		case 1:
			h = m - 1
		}
	}
	return -1
}
