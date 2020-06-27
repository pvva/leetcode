package main

func main() {
	println(isPowerOfTwo(1), true)
	println(isPowerOfTwo(16), true)
	println(isPowerOfTwo(218), false)
}

/*
https://leetcode.com/problems/power-of-two/

Given an integer, write a function to determine if it is a power of two.

Example 1:

Input: 1
Output: true
Explanation: 20 = 1
Example 2:

Input: 16
Output: true
Explanation: 24 = 16
Example 3:

Input: 218
Output: false
*/

// 2^n contains one 1 and other zeroes to the right
// 2^n - 1 contains all ones
// 2^n & (2^n - 1) should be zero, so just test n & (n-1)
// edge case: 0 is not a power of two
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
