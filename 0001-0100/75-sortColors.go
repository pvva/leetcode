package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	fmt.Println(nums)
	sortColors(nums)
	fmt.Println(nums)

	nums = []int{2, 0, 2, 1, 1, 0}
	fmt.Println(nums)
	sortColorsOnePass(nums)
	fmt.Println(nums)
}

/*
https://leetcode.com/problems/sort-colors/

Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color\
are adjacent, with the colors in the order red, white and blue.
Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.
Note: You are not suppose to use the library's sort function for this problem.

Example:

Input: [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

Follow up:

A rather straight forward solution is a two-pass algorithm using counting sort.
First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's
and followed by 2's.
Could you come up with a one-pass algorithm using only constant space?
*/

func sortColors(nums []int) {
	reds := 0
	whites := 0
	blues := 0
	for _, v := range nums {
		if v == 0 {
			reds++
		} else if v == 1 {
			whites++
		} else {
			blues++
		}
	}

	i := 0
	for ; reds > 0; i++ {
		nums[i] = 0
		reds--
	}
	for ; whites > 0; i++ {
		nums[i] = 1
		whites--
	}
	for ; blues > 0; i++ {
		nums[i] = 2
		blues--
	}
}

func sortColorsOnePass(nums []int) {
	zp := 0
	tp := len(nums) - 1
	cp := 0

	for cp <= tp {
		if nums[cp] == 0 {
			nums[cp], nums[zp] = nums[zp], nums[cp]
			zp++
			cp++
		} else if nums[cp] == 1 {
			cp++
		} else if nums[cp] == 2 {
			nums[cp], nums[tp] = nums[tp], nums[cp]
			tp--
		}
	}
}
