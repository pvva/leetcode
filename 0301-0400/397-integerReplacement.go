package main

func main() {
	println(integerReplacement(8), 3)
	println(integerReplacement(7), 4)
}

/*
https://leetcode.com/problems/integer-replacement/

Given a positive integer n and you can do operations as follow:

If n is even, replace n with n/2.
If n is odd, you can replace n with either n + 1 or n - 1.
What is the minimum number of replacements needed for n to become 1?

Example 1:

Input:
8

Output:
3

Explanation:
8 -> 4 -> 2 -> 1
*/

func integerReplacement(n int) int {
	cnt := 0

	for n > 1 {
		cnt++
		if n&1 == 0 {
			n >>= 1
		} else if n == 3 || ((n-1)>>1)&1 == 0 {
			n--
		} else {
			n++
		}
	}

	return cnt
}
