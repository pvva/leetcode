package main

import "fmt"

func main() {
	//fmt.Println(maxSubArrayBF([]int{4, 1, 2, 1}))
	fmt.Println(maxSubArrayBF([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArrayLinear([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the
largest sum and return its sum.

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
*/

func maxSubArrayBF(nums []int) (int, []int) {
	start := 0
	end := 0
	s := nums[0]
	t := 0

	for i := 0; i < len(nums); i++ {
		t = nums[i]
		for j := i + 1; j < len(nums); j++ {
			t += nums[j]
			if s < t {
				s = t
				start = i
				end = j
			}
		}
	}

	return s, nums[start : end+1]
}

// Kadane's algorithm
func maxSubArrayLinear(nums []int) (int, []int) {
	start := 0
	end := 0
	bestSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		t := currentSum + nums[i]
		if t < nums[i] {
			currentSum = nums[i]
			start = i
		} else {
			currentSum = t
		}
		if bestSum < currentSum {
			bestSum = currentSum
			end = i
		}
	}

	return bestSum, nums[start : end+1]
}
