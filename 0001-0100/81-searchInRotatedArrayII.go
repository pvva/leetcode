package main

func main() {
	println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0), true)
	println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3), false)
	println(search([]int{2, 5, 6, 0, 0, 1, 2}, 5), true)
	println(search([]int{2, 5, 6, 0, 0, 1, 2}, 6), true)
	println(search([]int{2, 2, 2, 2, 2, 1, 2}, 1), true)
	println(search([]int{2, 1, 2, 2, 2, 2, 2}, 1), true)
	println(search([]int{2, 1, 2, 2, 2, 2, 2}, 3), false)
	println(search([]int{3, 5, 1}, 3), true)

	println("--")
	println(search2([]int{2, 5, 6, 0, 0, 1, 2}, 0), true)
	println(search2([]int{2, 5, 6, 0, 0, 1, 2}, 3), false)
	println(search2([]int{2, 5, 6, 0, 0, 1, 2}, 5), true)
	println(search2([]int{2, 5, 6, 0, 0, 1, 2}, 6), true)
	println(search2([]int{2, 2, 2, 2, 2, 1, 2}, 1), true)
	println(search2([]int{2, 1, 2, 2, 2, 2, 2}, 1), true)
	println(search2([]int{2, 1, 2, 2, 2, 2, 2}, 3), false)
	println(search2([]int{3, 5, 1}, 3), true)
}

/*
https://leetcode.com/problems/search-in-rotated-sorted-array-ii/

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).

You are given a target value to search. If found in the array return true, otherwise return false.

Example 1:

Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true
Example 2:

Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false
Follow up:

This is a follow up problem to Search in Rotated Sorted Array, where nums may contain duplicates.
Would this affect the run-time complexity? How and why?
*/

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return true
		} else if nums[m] < nums[l] && nums[m] < nums[r] {
			// pivot is between l and m
			if target > nums[m] && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else if nums[m] > nums[l] && nums[m] > nums[r] {
			// pivot is between m and r
			if target < nums[m] && target >= nums[l] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else if nums[m] > nums[l] && nums[m] < nums[r] {
			// no pivot
			if target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if nums[l] == target {
				return true
			}
			l++
		}
	}

	return false
}

func search2(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return true
		}
		if nums[l] < nums[m] || nums[m] > nums[r] {
			if target > nums[m] || target < nums[l] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else if nums[m] < nums[r] || nums[m] < nums[l] {
			if target > nums[r] || target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			r--
		}
	}
	return false
}
