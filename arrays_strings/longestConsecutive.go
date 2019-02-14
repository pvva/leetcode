package main

import "sort"

func main() {
	//println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	//println(longestConsecutive([]int{}))
	//println(longestConsecutive([]int{0, -1}))
	println(longestConsecutive([]int{1, 2, 0, 1}))
}

/*
Given an unsorted array of integers, find the length of the longest consecutive elements sequence.
Your algorithm should run in O(n) complexity.
*/
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	c0 := 1
	c1 := 1
	v := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i] == v+1 {
			c0++
		} else {
			if c0 > c1 {
				c1 = c0
			}
			c0 = 1
		}
		v = nums[i]
	}

	if c1 > c0 {
		return c1
	}

	return c0
}
