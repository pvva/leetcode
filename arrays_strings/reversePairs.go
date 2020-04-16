package main

import "sort"

func main() {
	println(reversePairs([]int{1, 3, 2, 3, 1}), 2)
	println(reversePairs([]int{2, 4, 3, 5, 1}), 3)
	println(reversePairs([]int{1, 3, 2, 5, 6, 8}), 0)
	println(reversePairs([]int{-5, -5}), 1)
}

/*
https://leetcode.com/problems/reverse-pairs/

Given an array nums, we call (i, j) an important reverse pair if i < j and nums[i] > 2*nums[j].

You need to return the number of important reverse pairs in the given array.

Example1:

Input: [1,3,2,3,1]
Output: 2
Example2:

Input: [2,4,3,5,1]
Output: 3
Note:
The length of the given array will not exceed 50,000.
All the numbers in the input array are in the range of 32-bit integer.
*/

func reversePairs(nums []int) int {
	// O(n*log(n))
	bit := make([]int, len(nums)+1)

	copyNums := make([]int, len(nums))
	copy(copyNums, nums)

	sort.SliceStable(copyNums, func(i, j int) bool {
		return copyNums[i] < copyNums[j]
	})

	res := 0

	for _, v := range nums {
		res += search(bit, sort.SearchInts(copyNums, v*2+1)+1)
		insert(bit, sort.SearchInts(copyNums, v)+1)
	}

	return res
}

func search(bit []int, idx int) int {
	sum := 0

	i := idx
	for i < len(bit) {
		sum += bit[i]
		i += i & -i
	}

	return sum
}

func insert(bit []int, idx int) {
	i := idx
	for i > 0 {
		bit[i] += 1
		i -= i & -i
	}
}
