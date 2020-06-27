package main

func main() {
	println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}), 10)
	println(largestRectangleArea([]int{2, 1, 5, 6, 4, 3}), 12)
}

/*
https://leetcode.com/problems/largest-rectangle-in-histogram/

Given n non-negative integers representing the histogram's bar height where the width of each bar is 1,
find the area of largest rectangle in the histogram.

Example:

Input: [2,1,5,6,2,3]
Output: 10
*/

// the idea is to use "next lesser element" approach (NLE), which is calculated using monotonous increasing stack
// (in contrary, if you want next greater element, use monotonous decreasing stack)
// when going from left to right, then for each index i NLE to the left is calculated
// when going from right to left, then for each index i NLE to the right is calculated
// at this point let's take a look at some index i in the original array:
//    using NLE from left and NLE from right we can say that max square at this index would be right NLE index minus
//    left NLE index minus 1 (since we are effectively skipping values that stand on those indices, as they are less then
//    current value) and multiplied by current value at index i, because it is obvious that all other items in array
//    to the left that come before left NLE and the same to the right are more or equal to the current value, so current value
//    is the height if the rectangular, while the indices difference is the width
// with this observation we just need to go over an array and calculate the max possible square at every point and select
// the maximum one
func largestRectangleArea(heights []int) int {
	sq := 0

	mins := make([][2]int, len(heights))   // left, right
	stacks := make([][2]int, len(heights)) // left, right

	lg := len(heights)
	il := -1
	ir := -1

	for i, v := range heights {
		for il >= 0 && heights[stacks[il][0]] >= v {
			il--
		}

		if il < 0 {
			mins[i][0] = -1
		} else {
			mins[i][0] = stacks[il][0]
		}
		il++
		stacks[il][0] = i

		for ir >= 0 && heights[stacks[ir][1]] >= heights[lg-1-i] {
			ir--
		}
		if ir < 0 {
			mins[lg-i-1][1] = lg
		} else {
			mins[lg-i-1][1] = stacks[ir][1]
		}

		ir++
		stacks[ir][1] = lg - i - 1
	}

	for i, v := range heights {
		s := (mins[i][1] - mins[i][0] - 1) * v
		if s > sq {
			sq = s
		}
	}

	return sq
}
