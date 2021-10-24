package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("[[1,1,2],[1,2,1],[2,1,1]]")
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println("[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]")
	fmt.Println(permuteUnique([]int{1, 2, 3}))
}

/*
https://leetcode.com/problems/permutations-ii/

Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.

Example 1:

Input: nums = [1,1,2]
Output:
[[1,1,2],
 [1,2,1],
 [2,1,1]]

Example 2:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Constraints:

1 <= nums.length <= 8
-10 <= nums[i] <= 10
*/

func permute2(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	if len(nums) == 2 {
		if nums[0] == nums[1] {
			return [][]int{nums}
		}
		return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	}
	res := [][]int{}

	for i := 0; i < len(nums); i++ {
		for ; i < len(nums)-1 && nums[i] == nums[i+1]; i++ {
		}

		t := []int{}
		t = append(t, nums[:i]...)
		t = append(t, nums[i+1:]...)
		r := permute2(t)
		for _, p := range r {
			res = append(res, append(p, nums[i]))
		}
	}

	return res
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	return permute2(nums)
}
