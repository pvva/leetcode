package main

import (
	"math"
	"sort"
)

func main() {
	println(2, threeSumClosest([]int{-1, 2, 1, -4}, 1))
	println(0, threeSumClosest([]int{0, 0, 0}, 1))
	println(3, threeSumClosest([]int{1, 1, -1, -1, 3}, 3))
	println(3, threeSumClosest([]int{1, 1, 1, 1, 1, 1, 1}, 3))
}

/*
https://leetcode.com/problems/3sum-closest/

Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.

Return the sum of the three integers.

You may assume that each input would have exactly one solution.


Example 1:

Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

Example 2:

Input: nums = [0,0,0], target = 1
Output: 0

Constraints:

3 <= nums.length <= 1000
-1000 <= nums[i] <= 1000
-10^4 <= target <= 10^4
*/

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	l := len(nums) - 1
	result := 0
	diff := math.MaxInt32
	for i := 0; i < l-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		secondIdx := i + 1
		thirdIdx := l
		for secondIdx < thirdIdx {
			sum := nums[i] + nums[secondIdx] + nums[thirdIdx]
			d := sum - target
			if d == 0 {
				return sum
			} else if d < 0 {
				secondIdx++
			} else {
				thirdIdx--
			}

			if d < 0 {
				d = -d
			}
			if d < diff {
				diff = d
				result = sum
			}
		}
	}

	return result
}
