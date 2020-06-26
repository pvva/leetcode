package main

func main() {
	println(findDuplicate([]int{1, 3, 4, 2, 2}))
	println(findDuplicate([]int{3, 1, 3, 4, 2}))
	println(findDuplicate([]int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}))
}

/*
https://leetcode.com/problems/find-the-duplicate-number/


Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive),
prove that at least one duplicate number must exist. Assume that there is only one duplicate number,
find the duplicate one.

Example 1:

Input: [1,3,4,2,2]
Output: 2

Example 2:

Input: [3,1,3,4,2]
Output: 3

Note:

You must not modify the array (assume the array is read only).
You must use only constant, O(1) extra space.
Your runtime complexity should be less than O(n2).
There is only one duplicate number in the array, but it could be repeated more than once.
*/

func findDuplicate(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	s := nums[0]
	f := nums[s]

	for s != f {
		s = nums[s]
		f = nums[nums[f]]
	}

	f = 0
	for s != f {
		f = nums[f]
		s = nums[s]
	}

	return s
}
