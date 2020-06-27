package main

import (
	"sort"
)

func main() {
	println(firstMissingPositiveBF([]int{1, 2, 0}))
	println(firstMissingPositiveBF([]int{3, 4, -1, 1}))
	println(firstMissingPositiveBF([]int{7, 8, 9, 11, 12}))

	println(firstMissingPositive([]int{1, 2, 0}))
	println(firstMissingPositive([]int{3, 4, -1, 1}))
	println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
	println(firstMissingPositive([]int{1}))
}

/*
https://leetcode.com/problems/first-missing-positive/

Given an unsorted integer array, find the smallest missing positive integer.

Example 1:

Input: [1,2,0]
Output: 3
Example 2:

Input: [3,4,-1,1]
Output: 2
Example 3:

Input: [7,8,9,11,12]
Output: 1
Note:

Your algorithm should run in O(n) time and uses constant extra space.
*/

func firstMissingPositiveBF(nums []int) int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	dv := 1
	for _, v := range nums {
		if v <= 0 {
			continue
		}

		if v == dv {
			dv++
		} else if v > dv {
			break
		}
	}

	return dv
}

func firstMissingPositive(nums []int) int {
	l := len(nums)

	for i, v := range nums {
		if v <= 0 {
			nums[i] = l + 1
		}
	}
	for _, v := range nums {
		av := v
		if av < 0 {
			av = -av
		}

		if av <= l && nums[av-1] > 0 {
			nums[av-1] *= -1
		}
	}
	for i, v := range nums {
		if v < 0 {
			continue
		}

		return i + 1
	}

	return l + 1
}
