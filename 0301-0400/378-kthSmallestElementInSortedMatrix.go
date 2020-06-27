package main

import (
	"sort"
)

func main() {
	println(kthSmallest([][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 14, 15},
	}, 8))
	println(kthSmallest([][]int{
		{1, 3, 5},
		{6, 7, 12},
		{11, 14, 14},
	}, 3))
	println(kthSmallest([][]int{
		{1, 2, 2, 4},
		{3, 4, 6, 7},
		{4, 5, 7, 8},
		{6, 7, 9, 10},
	}, 8)) //3, 4, 7, 8, 9, 11, 14
	println(kthSmallest([][]int{
		{3, 6, 9, 12, 17, 22},
		{5, 11, 11, 16, 22, 26},
		{10, 11, 14, 16, 24, 31},
		{10, 15, 17, 17, 29, 31},
		{14, 17, 20, 23, 34, 37},
		{19, 21, 22, 28, 37, 40},
	}, 35))

	println(kthSmallest([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 20))
}

/*
https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/

Given a n x n matrix where each of the rows and columns are sorted in ascending order, find the kth smallest
element in the matrix.
Note that it is the kth smallest element in the sorted order, not the kth distinct element.

Example:

matrix = [
   [ 1,  5,  9],
   [10, 11, 13],
   [12, 13, 15]
],
k = 8,

return 13.

Note:
You may assume k is always valid, 1 ≤ k ≤ n2.*/

func kthSmallest(matrix [][]int, k int) int {
	return kthSmallestBF(matrix, k)
}

func kthSmallestBF(matrix [][]int, k int) int {
	n := len(matrix)

	if n == 1 {
		return matrix[0][0]
	}

	flat := make([]int, n*n)

	c := 0
	for _, mx := range matrix {
		for _, v := range mx {
			flat[c] = v
			c++
		}
	}

	sort.SliceStable(flat, func(i, j int) bool {
		return flat[i] < flat[j]
	})

	return flat[k-1]
}
