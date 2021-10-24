package main

import "fmt"

func main() {
	fmt.Println("[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]")
	fmt.Println(permute([]int{1, 2, 3}))
}

/*
https://leetcode.com/problems/permutations/

Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:

Input: nums = [0,1]
Output: [[0,1],[1,0]]

Example 3:

Input: nums = [1]
Output: [[1]]

Constraints:

1 <= nums.length <= 6
-10 <= nums[i] <= 10
All the integers of nums are unique.
*/

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	if len(nums) == 2 {
		return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	}
	res := [][]int{}

	for i, v := range nums {
		t := []int{}
		t = append(t, nums[:i]...)
		t = append(t, nums[i+1:]...)
		r := permute(t)
		for _, p := range r {
			res = append(res, append(p, v))
		}
	}

	return res
}
