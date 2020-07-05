package main

func main() {
	println(hammingDistance(1, 4), 2)
}

/*
https://leetcode.com/problems/hamming-distance/

The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Given two integers x and y, calculate the Hamming distance.

Note:
0 ≤ x, y < 231.

Example:

Input: x = 1, y = 4

Output: 2

Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑

The above arrows point to positions where the corresponding bits are different.
*/

func hammingDistance(x int, y int) int {
	c := 0

	for x > 0 || y > 0 {
		c += x&1 ^ y&1
		x >>= 1
		y >>= 1
	}

	return c
}

func hammingDistance2(x int, y int) int {
	c := 0
	t := x ^ y

	for t > 0 {
		c, t = c+t&1, t>>1
	}

	return c
}

func hammingDistance3(x int, y int) int {
	c := 0
	t := x ^ y

	for t > 0 {
		t, c = t&(t-1), c+1
	}

	return c
}
