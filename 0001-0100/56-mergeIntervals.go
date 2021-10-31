package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("[[1,6],[8,10],[15,18]]", merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println("[[1,4]]", merge([][]int{{1, 4}, {2, 3}}))
}

/*
https://leetcode.com/problems/merge-intervals/

Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return
an array of the non-overlapping intervals that cover all the intervals in the input.

Example 1:

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].

Example 2:

Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.

Constraints:

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104
*/

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}

	for i := 0; i < len(intervals); {
		min := intervals[i][0]
		max := intervals[i][1]
		for ; i < len(intervals) && intervals[i][0] <= max; i++ {
			if max < intervals[i][1] {
				max = intervals[i][1]
			}
		}
		res = append(res, []int{min, max})
	}

	return res
}
