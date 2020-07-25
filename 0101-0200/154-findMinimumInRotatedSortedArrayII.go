package main

func main() {
	println(findMin([]int{1, 3, 5}), 1)
	println(findMin([]int{2, 2, 2, 0, 1}), 0)
	println(findMin([]int{1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1}), 0)
}

/*
https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).

Find the minimum element.

The array may contain duplicates.

Example 1:

Input: [1,3,5]
Output: 1

Example 2:

Input: [2,2,2,0,1]
Output: 0

Note:

This is a follow up problem to Find Minimum in Rotated Sorted Array.
Would allow duplicates affect the run-time complexity? How and why?
*/

// this should not be "hard" level problem
func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		m := l + (r-l)>>1
		if nums[m] < nums[r] {
			r = m
		} else if nums[m] > nums[r] {
			l = m + 1
		} else {
			r--
		}
	}

	return nums[r]
}
