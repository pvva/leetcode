package main

func main() {
	println(rangeBitwiseAnd(5, 7), 4)
	println(rangeBitwiseAnd(0, 1), 0)
}

/*
https://leetcode.com/problems/bitwise-and-of-numbers-range/

Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND of all numbers in this range, inclusive.

Example 1:

Input: [5,7]
Output: 4

Example 2:

Input: [0,1]
Output: 0
*/

func rangeBitwiseAnd(m int, n int) int {
	shiftCounter := 0

	// shift right until both number are equal, which means that they have common prefix, which is a bases for an answer
	// to get an answer we need to shift the basis left as many times as we shifted original number to the right
	for m != n {
		m >>= 1
		n >>= 1
		shiftCounter++
	}

	return n << shiftCounter
}
