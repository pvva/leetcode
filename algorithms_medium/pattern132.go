package main

import "math"

func main() {
	println(find132patternLinear([]int{1, 2, 3, 4}), false)
	println(find132patternLinear([]int{3, 1, 4, 2}), true)
	println(find132patternLinear([]int{-1, 3, 2, 0}), true)
	println(find132patternLinear([]int{1, 0, 1, -4, -3}), false)
	println(find132patternLinear([]int{-2, 1, 2, -2, 1, 2}), true)
	println(find132patternLinear([]int{3, 5, 0, 3, 4}), true)
	println(find132patternLinear([]int{3, 5, 0, 3, 3}), false)
}

/*
https://leetcode.com/problems/132-pattern/

Given a sequence of n integers a1, a2, ..., an, a 132 pattern is a subsequence ai, aj, ak such that i < j < k
and ai < ak < aj. Design an algorithm that takes a list of n numbers as input and checks whether there is a
132 pattern in the list.

Note: n will be less than 15,000.

Example 1:
Input: [1, 2, 3, 4]

Output: False

Explanation: There is no 132 pattern in the sequence.

Example 2:
Input: [3, 1, 4, 2]

Output: True

Explanation: There is a 132 pattern in the sequence: [1, 4, 2].

Example 3:
Input: [-1, 3, 2, 0]

Output: True

Explanation: There are three 132 patterns in the sequence: [-1, 3, 2], [-1, 3, 0] and [-1, 2, 0].
*/

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	ai := nums[0]
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] <= ai {
			ai = nums[i]
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] && nums[j] > ai {
				return true
			}
		}

	}

	return false
}

func find132patternLinear(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	tStack := make([]int, len(nums))
	tIdx := -1

	ak := math.MinInt32
	for i := len(nums) - 1; i >= 0; i-- {
		// nums[i] is supposed ai
		// ak is previously calculated and aj sits on top of the stack, so if stack is not empty and
		//   supposed ai is less then ak, then the pattern is found
		if nums[i] < ak && tIdx >= 0 {
			return true
		}
		// nums[i] failed to be ai, now treat it as possible aj

		// need to find ak such that ak < aj
		// due to index ordering (top to bottom) and using stack (it reverses the order)
		//   it is guaranteed that k > j if we find in stack an element such that it is less
		//   than current supposed aj and is as deep in the stack as possible

		for tIdx >= 0 && tStack[tIdx] < nums[i] {
			// need to skip elements that are less then currently supposed aj as much as possible
			// so that ak has largest index (stands as much to the right as possible)
			ak = tStack[tIdx]
			tIdx--
		}

		// at this point ak is less then current nums[i] (supposed aj) and is furthest possible to the right in the array

		// supposed aj goes to the top of the stack
		tIdx++
		tStack[tIdx] = nums[i]
	}

	return false
}
