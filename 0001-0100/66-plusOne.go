package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}), []int{1, 2, 4})
	fmt.Println(plusOne([]int{4, 3, 2, 1}), []int{4, 3, 2, 2})
	fmt.Println(plusOne([]int{1, 9}), []int{2, 0})
	fmt.Println(plusOne([]int{9, 9, 9}), []int{1, 0, 0, 0})
}

/*
https://leetcode.com/problems/plus-one/

Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array
contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
*/

func plusOne(digits []int) []int {
	carry := 0
	l := len(digits) - 1
	digits[l]++

	for i := l; i >= 0; i-- {
		digits[i] += carry
		carry = 0

		if digits[i] == 10 {
			digits[i] = 0
			carry = 1
		} else {
			break
		}
	}
	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}

	return digits
}
