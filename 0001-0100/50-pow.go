package main

func main() {
	println(myPow(2, 2147483648))
	println(myPow(2, -2147483648))
	println(myPow(2, -3))
	println(myPow(2, 7))
	println(myPow(2, 4))
}

/*
https://leetcode.com/problems/powx-n/

Implement pow(x, n), which calculates x raised to the power n (xn).

Example 1:

Input: 2.00000, 10
Output: 1024.00000

Example 2:

Input: 2.10000, 3
Output: 9.26100

Example 3:

Input: 2.00000, -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
Note:

-100.0 < x < 100.0
n is a 32-bit signed integer, within the range [−231, 231 − 1]
*/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	neg := false
	if n < 0 {
		neg = true
		n = -n
	}

	c := 1.0
	if n&1 == 1 {
		c = x
	}
	r := c * myPow(x*x, n>>1)
	if neg {
		return 1 / r
	}

	return r
}

func myPowIterative(x float64, n int) float64 {
	r := 1.0

	neg := false
	if n < 0 {
		neg = true
		n = -n
	}

	for n > 0 {
		if n&1 == 1 {
			r *= x
		}
		x *= x
		n >>= 1
	}

	if neg {
		return 1 / r
	}

	return r
}
