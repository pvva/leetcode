package main

func main() {
	println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), 6)
	println(trap2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), 6)
}

/*
https://leetcode.com/problems/trapping-rain-water/

Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much
water it is able to trap after raining.

Example:

Input: [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
*/

func trap(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	maxSoFar := make([][2]int, len(height)) // values: left, right

	lg := len(height) - 1
	maxSoFar[0][0] = height[0]
	maxSoFar[lg][1] = height[lg]

	for i := 1; i < lg; i++ {
		maxSoFar[i][0] = max(maxSoFar[i-1][0], height[i])
		maxSoFar[lg-i][1] = max(maxSoFar[lg-i+1][1], height[lg-i])
	}
	maxSoFar[lg][0] = max(maxSoFar[lg-1][0], height[lg])
	maxSoFar[0][1] = max(maxSoFar[1][1], height[0])

	sum := 0

	for i, v := range height {
		sum += min(maxSoFar[i][0], maxSoFar[i][1]) - v
	}

	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// incredibly amazing idea with just two pointers (not mine)
func trap2(height []int) int {
	li := 0
	ri := len(height) - 1

	lmax := 0
	rmax := 0

	total := 0

	for li < ri {
		lmax = max(lmax, height[li])
		rmax = max(rmax, height[ri])

		if lmax < rmax {
			total += lmax - height[li]
			li++
		} else {
			total += rmax - height[ri]
			ri--
		}
	}

	return total
}
