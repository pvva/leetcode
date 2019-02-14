package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSlidingWindowBF([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindowBF([]int{1, -1}, 1))

	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1, -1}, 1))
}

/*
Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right.
You can only see the k numbers in the window. Each time the sliding window moves right by one position. Return the max sliding window.
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
Output: [3,3,5,5,6,7]
*/

func maxSlidingWindowBF(nums []int, k int) []int {
	l := len(nums) - k
	if len(nums) == 0 || l < 0 {
		return []int{}
	}
	result := make([]int, l+1)

	for i := 0; i <= l; i++ {
		result[i] = max(nums[i : i+k])
	}

	return result
}

func max(a []int) int {
	m := math.MinInt32
	for _, v := range a {
		if m < v {
			m = v
		}
	}

	return m
}

func maxSlidingWindow(nums []int, k int) []int {
	l := len(nums) - k
	if len(nums) == 0 || l < 0 {
		return []int{}
	}
	result := make([]int, l+1)
	if k == 1 {
		result[0] = nums[0]
	}

	window := []int{0}
	for i := 1; i < len(nums); i++ {
		if i-k >= window[0] {
			window = window[1:]
		}

		l = len(window) - 1
		for l >= 0 && nums[window[l]] < nums[i] {
			window = window[:l]
			l--
		}
		window = append(window, i)

		if i >= k-1 {
			result[i-k+1] = nums[window[0]]
		}
	}

	return result
}
