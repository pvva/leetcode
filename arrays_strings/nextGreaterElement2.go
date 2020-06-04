package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}

/*
https://leetcode.com/problems/next-greater-element-ii/

Given a circular array (the next element of the last element is the first element of the array), print the
Next Greater Number for every element. The Next Greater Number of a number x is the first greater number to
its traversing-order next in the array, which means you could search circularly to find its next greater number.
If it doesn't exist, output -1 for this number.

Example 1:
Input: [1,2,1]
Output: [2,-1,2]
Explanation: The first 1's next greater number is 2;
The number 2 can't find next greater number;
The second 1's next greater number needs to search circularly, which is also 2.
*/

// the idea is to use monotonous increasing stack, keep values in stack
func nextGreaterElements(nums []int) []int {
	stack := make([]int, len(nums))
	si := -1
	result := make([]int, len(nums))

	for i := 2*len(nums) - 2; i >= 0; i-- {
		idx := i % len(nums)
		for si >= 0 && stack[si] <= nums[idx] {
			si--
		}
		if si >= 0 {
			result[idx] = stack[si]
		} else {
			result[idx] = -1
		}

		si++
		stack[si] = nums[idx]
	}

	return result
}
