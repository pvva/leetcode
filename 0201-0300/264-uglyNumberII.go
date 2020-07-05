package main

func main() {
	println(nthUglyNumber(1), 1)
	println(nthUglyNumber(10), 12)
	println(nthUglyNumber(11), 15)
	println(nthUglyNumber(12), 16)
}

/*
https://leetcode.com/problems/ugly-number-ii/

Write a program to find the n-th ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.

Example:

Input: n = 10
Output: 12
Explanation: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10 ugly numbers.
Note:

1 is typically treated as an ugly number.
n does not exceed 1690.
*/

func nthUglyNumber(n int) int {
	numbers := make([]int, n)
	numbers[0] = 1
	c2, c3, c5 := 0, 0, 0

	for i := 1; i < n; i++ {
		n2, n3, n5 := numbers[c2]*2, numbers[c3]*3, numbers[c5]*5
		numbers[i] = min3(n2, n3, n5)

		if numbers[i] == n2 {
			c2++
		}
		if numbers[i] == n3 {
			c3++
		}
		if numbers[i] == n5 {
			c5++
		}
	}

	return numbers[n-1]
}

func min3(a, b, c int) int {
	return min2(min2(a, b), c)
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
