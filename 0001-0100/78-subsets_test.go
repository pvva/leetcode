package main

import (
	"testing"
)

func subsetsBitMask(nums []int) [][]int {
	result := [][]int{{}}

	mask := make([]bool, len(nums))
	for nextCombinationForMask(mask) {
		result = append(result, getSubsetByMask(nums, mask))
	}

	return result
}

func nextCombinationForMask(mask []bool) bool {
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

func getSubsetByMask(nums []int, mask []bool) []int {
	s := []int{}

	for i, v := range nums {
		if mask[i] {
			s = append(s, v)
		}
	}

	return s
}

// backtracking approach
func subsetsRecursiveBacktracking(nums []int) [][]int {
	subsets := [][]int{}
	generateSubsetRec(0, nums, []int{}, &subsets)

	return subsets
}

func generateSubsetRec(pos int, nums, subset []int, subsets *[][]int) {
	if pos == len(nums) {
		*subsets = append(*subsets, subset)
		return
	}

	skipSet := make([]int, len(subset)) // skipping current element
	copy(skipSet, subset)
	subset = append(subset, nums[pos]) // putting current element in

	generateSubsetRec(pos+1, nums, subset, subsets)  // left branch - include elements
	generateSubsetRec(pos+1, nums, skipSet, subsets) // right branch - exclude elements
}

var subset [][]int
var originalSets = [][]int{
	{1, 2, 3},                // 8
	{1, 2, 3, 4, 5},          // 32
	{1, 2, 3, 4, 5, 6, 7, 8}, // 256
}

func Benchmark_subsetsBitMask(b *testing.B) {
	var s [][]int
	for i := 0; i < b.N; i++ {
		for _, set := range originalSets {
			subsetsBitMask(set)
		}
	}

	subset = s
}

func Benchmark_subsetsRecursiveBacktracking(b *testing.B) {
	var s [][]int
	for i := 0; i < b.N; i++ {
		for _, set := range originalSets {
			subsetsRecursiveBacktracking(set)
		}
	}

	subset = s
}
