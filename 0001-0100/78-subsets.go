package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets2([]int{1, 2, 3}))
}

/*
https://leetcode.com/problems/subsets/

Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

// base idea: set power is 2^n, so we can use bitmask to define what elements to be included in the
// current subset
// since number of elements in original set can be more than 64, I'm using []bool as bitmask
func subsets(nums []int) [][]int {
	result := [][]int{{}}

	mask := make([]bool, len(nums))
	for nextCombination(mask) {
		result = append(result, getSubset(nums, mask))
	}

	return result
}

func nextCombination(mask []bool) bool {
	for i := 0; i < len(mask); i++ {
		if mask[i] {
			mask[i] = false
		} else {
			mask[i] = true
			return true
		}
	}

	return false
}

func getSubset(nums []int, mask []bool) []int {
	s := []int{}

	for i, v := range nums {
		if mask[i] {
			s = append(s, v)
		}
	}

	return s
}

// backtracking approach, faster than bitmasking
func subsets2(nums []int) [][]int {
	subsets := [][]int{}
	generateSubset(0, nums, []int{}, &subsets)

	return subsets
}

func generateSubset(pos int, nums, subset []int, subsets *[][]int) {
	if pos == len(nums) {
		*subsets = append(*subsets, subset)
		return
	}

	skipSet := make([]int, len(subset)) // skipping current element
	copy(skipSet, subset)
	subset = append(subset, nums[pos]) // putting current element in

	generateSubset(pos+1, nums, subset, subsets)  // left branch - include elements
	generateSubset(pos+1, nums, skipSet, subsets) // right branch - exclude elements
}
