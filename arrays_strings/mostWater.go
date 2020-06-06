package main

func main() {
	println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	println(maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

	println(maxArea([]int{1, 2, 1}))
	println(maxArea2([]int{1, 2, 1}))

	println(maxArea([]int{1, 2, 4, 3}))
	println(maxArea2([]int{1, 2, 4, 3}))

	println(maxArea([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	println(maxArea2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

/*
https://leetcode.com/problems/container-with-most-water/

Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines,
which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.

Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49
*/

func maxArea(height []int) int {
	area := 0

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			c := min(height[i], height[j]) * (j - i)
			if area < c {
				area = c
			}
		}
	}

	return area
}

func maxArea2(height []int) int {
	area := 0
	i := 0
	e := len(height) - 1

	for i < e {
		a := min(height[i], height[e]) * (e - i)
		if a > area {
			area = a
		}
		if height[i] < height[e] {
			i++
		} else {
			e--
		}
	}

	return area
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
