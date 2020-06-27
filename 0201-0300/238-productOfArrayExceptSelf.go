package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{2, 4, 6, 8}))
}

/*
https://leetcode.com/problems/product-of-array-except-self/

Given an array nums of n integers where n > 1,  return an array output such that output[i] is equal to the product of
all the elements of nums except nums[i].

Example:

Input:  [1,2,3,4]
Output: [24,12,8,6]
Constraint: It's guaranteed that the product of the elements of any prefix or suffix of the array
(including the whole array) fits in a 32 bit integer.

Note: Please solve it without division and in O(n).

Follow up:
Could you solve it with constant space complexity? (The output array does not count as extra space for the purpose
of space complexity analysis.)
*/

func productExceptSelf(nums []int) []int {
	l := len(nums)

	out := make([]int, l)

	// all products except for last element
	out[0] = 1
	for i := 1; i < l; i++ {
		out[i] = out[i-1] * nums[i-1]
	}

	// all products backward with skipping current element
	product := nums[l-1]
	for i := l - 2; i >= 0; i-- {
		out[i] *= product
		product *= nums[i]
	}

	return out
}
