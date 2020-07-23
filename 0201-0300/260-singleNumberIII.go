package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}), []int{3, 5})
}

/*
https://leetcode.com/problems/single-number-iii/

Given an array of numbers nums, in which exactly two elements appear only once and all the other elements appear
exactly twice. Find the two elements that appear only once.

Example:

Input:  [1,2,1,3,2,5]
Output: [3,5]
Note:

The order of the result is not important. So in the above example, [5, 3] is also correct.
Your algorithm should run in linear runtime complexity. Could you implement it using only constant space complexity?
*/

func singleNumber(nums []int) []int {
	xor2 := 0

	for _, v := range nums {
		xor2 ^= v
	}

	lsb := 0
	for i := 0; lsb == 0; i++ {
		lsb = xor2 & (1 << i)
	}

	n1, n2 := 0, 0
	for _, v := range nums {
		if v&lsb > 0 {
			n1 ^= v
		} else {
			n2 ^= v
		}
	}

	return []int{n1, n2}
}
