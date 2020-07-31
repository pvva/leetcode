package main

import "math"

func main() {
	println(climbStairs(2), 2)
	println(climbStairs(3), 3)
	println(climbStairs(4), 5)
	println(climbStairs(45), 1836311903)

	println(climbStairsFibo(2), 2)
	println(climbStairsFibo(3), 3)
	println(climbStairsFibo(4), 5)
	println(climbStairsFibo(45), 1836311903)
}

/*
https://leetcode.com/problems/climbing-stairs/

You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
*/

func climbStairs(n int) int {
	return climbStairsWithMem(n, make([]int, n+1))
}

func climbStairsWithMem(n int, mem []int) int {
	if n < 3 {
		return n
	}
	if mem[n] > 0 {
		return mem[n]
	}
	mem[n] = climbStairsWithMem(n-1, mem) + climbStairsWithMem(n-2, mem)

	return mem[n]
}

func climbStairsFibo(n int) int {
	sq5 := math.Sqrt(5)
	phi := (1.0 + sq5) / 2.0
	// Binet's formula
	// Fn ~= (phi ^ n) / Sqrt(5)
	// since fibo numbers start with 1, 1 we should skip 1st, that's why we use n+1 as power value

	return int(math.Round(math.Pow(phi, float64(n+1)) / sq5))
}
