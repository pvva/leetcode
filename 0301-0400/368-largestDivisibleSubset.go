package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestDivisibleSubset([]int{1, 2, 4, 8}), []int{1, 2, 4, 8})
	fmt.Println(largestDivisibleSubset([]int{1, 2, 3}), []int{1, 2})
	fmt.Println(largestDivisibleSubset([]int{1, 2, 3, 4, 8}), []int{1, 2, 4, 8})
}

/*
https://leetcode.com/problems/largest-divisible-subset/

Given a set of distinct positive integers, find the largest subset such that every pair (Si, Sj) of elements
in this subset satisfies:

Si % Sj = 0 or Sj % Si = 0.

If there are multiple solutions, return any subset is fine.

Example 1:

Input: [1,2,3]
Output: [1,2] (of course, [1,3] will also be ok)

Example 2:

Input: [1,2,4,8]
Output: [1,2,4,8]
*/

// Let's start with the following observation: if at position j the number is divisible by number at position i,
//   than number at position j extends the set that ended at position i.
// Easiest way to track divisibility of number is to follow from lowest to largest number, so array should be sorted.
// Then let's keep track 2 things: the size of the set that starts at position i, and next index in input array that
//   points to the next number in the given set at position i.
// Iterate from largest to lowest number, and inside iterate back from current number to the largest to check inclusion
//   of numbers into different sets.
// Suppose that outer loop is at position i, and inner loop is at position j.
// If j == i, then current number forms the set of it's own with length 1.
// if j > i and N[j] % N[i] (since array is sorted it is guaranteed that N[j] > N[i]) then number at position j extends
//   the set that ended at position i. We might already know the length of the set that went through position j, so we need to
//   check if length at position i increased by 1 (L[i]+1) is larger than length at position j (L[j]). If it is so, than
//    we extend the set, incrementing the length at position i by 1 and setting position of next element in this new set to be
//    j (Pos[i] = j).
// Dynamic programming here comes from the fact that we track lengths of sets that went through some position.
func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)

	lengths := make([]int, len(nums))
	positions := make([]int, len(nums))

	maxLength := 0
	currentPos := 0

	for i := len(nums) - 1; i >= 0; i-- {
		for j := i; j < len(nums); j++ {
			if nums[j]%nums[i] == 0 && lengths[i] < lengths[j]+1 {
				lengths[i] = lengths[j] + 1
				positions[i] = j
			}
		}

		if maxLength < lengths[i] {
			maxLength = lengths[i]
			currentPos = i
		}
	}

	result := make([]int, maxLength)
	for i := 0; i < maxLength; i++ {
		result[i] = nums[currentPos]
		currentPos = positions[currentPos]
	}

	return result
}
