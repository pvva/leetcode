package main

import "math"

func main() {
	println(numSquares(12))
	println(numSquares(13))
	println(numSquares(163))
	println(numSquaresLagrange(163))
}

/*
Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.

Example 1:

Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.
Example 2:

Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.
*/

func numSquares(n int) int {
	return ns(n, make(map[int]int))
}

func ns(n int, minSq map[int]int) int {
	if n == 0 {
		return 0
	}
	if n < 4 {
		return n
	}

	sr := int(math.Floor(math.Sqrt(float64(n))))
	if n == sr*sr {
		minSq[n] = 1

		return 1
	}

	ms := n

	for sr > 1 {
		q := sr * sr
		c := 1
		r := n - q
		for r > 0 {
			r -= q
			c++
		}
		if r < 0 {
			r += q
			c--
		}

		v, ex := minSq[r]
		if !ex {
			v = ns(r, minSq)
			minSq[r] = v
		}
		ms = min(ms, v+c)

		sr--
	}

	return ms
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func numSquaresLagrange(n int) int {
	v := n
	for v&3 == 0 { // n%4 == 0
		v >>= 2
	}
	if v&7 == 7 { // n%7 == 7
		return 4
	}

	q := 0
	sq := 0

	for sq <= n {
		r := int(math.Floor(math.Sqrt(float64(n - sq))))

		if sq+r*r == n {
			v := 0
			if q > 0 {
				v++
			}
			if r > 0 {
				v++
			}
			return v
		}

		q++
		sq = q * q
	}

	return 3
}
