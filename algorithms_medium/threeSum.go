package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}), "---", [][]int{{-1, 0, 1}, {-1, -1, 2}})
	fmt.Println(threeSum([]int{3, 0, -2, -1, 1, 2}), "---", [][]int{{-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}})
	fmt.Println(threeSum([]int{-1, -1, 0, 0, 1, 1}), "---", [][]int{{-1, 0, 1}})
	fmt.Println(threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}), "---",
		[][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}})
}

/*
https://leetcode.com/problems/3sum/

Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0?
Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

func threeSum(nums []int) [][]int {
	result := [][]int{}

	sort.Ints(nums)
	l := len(nums) - 1
	for i := 0; i < l-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		secondIdx := i + 1
		thirdIdx := l
		for secondIdx < thirdIdx {
			sum := nums[i] + nums[secondIdx] + nums[thirdIdx]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[secondIdx], nums[thirdIdx]})
				secondIdx++
				thirdIdx--

				for secondIdx < thirdIdx && nums[secondIdx-1] == nums[secondIdx] {
					secondIdx++
				}
				for secondIdx < thirdIdx && nums[thirdIdx+1] == nums[thirdIdx] {
					thirdIdx--
				}
			} else if sum < 0 {
				secondIdx++
			} else {
				thirdIdx--
			}
		}
	}

	return result
}
